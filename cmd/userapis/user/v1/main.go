package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	gormdriver "github.com/tuanden0/refactor_user_ms/internal/databases/mysql/gorm_driver"
	logger "github.com/tuanden0/refactor_user_ms/internal/logs/zap_driver"
	"github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/repositories"
	userV1 "github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/services"
	userVD "github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/validators"
	vd "github.com/tuanden0/refactor_user_ms/internal/validators"
	userV1PB "github.com/tuanden0/refactor_user_ms/proto/gen/go/user/v1"
)

var (
	gRPCAddr string
	gwAddr   string
)

// Custom error struct to handle custom error return
type errorBody struct {
	Code    int32             `json:"code,omitempty"`
	Message string            `json:"message,omitempty"`
	Details map[string]string `json:"details,omitempty"`
}

func init() {
	flag.StringVar(&gRPCAddr, "grpc", "0.0.0.0:50001", "gRPC server address:port")
	flag.StringVar(&gwAddr, "gw", "0.0.0.0:8000", "gateway server address:port")
	flag.Parse()
}

func main() {

	// Connect to DB
	logger.Info("Init database", nil)
	gormDB, gormErr := gormdriver.ConnectDatabase()
	if gormErr != nil {
		logger.Error("database %v", gormErr)
		log.Fatal(gormErr)
	}
	logger.Info("Init database success", nil)

	// Create user repo manager
	logger.Info("Init user repository manager", nil)
	userRepo := repositories.NewManager(gormDB, logger.Log)
	logger.Info("Init user repository manager success", nil)

	// Create global validator
	logger.Info("Init validator", nil)
	validator := vd.NewValidator(*logger.Log)
	validator.InitValidate()
	if validatorErr := validator.InitTranslator(); validatorErr != nil {
		logger.Error("failed to init validator translator %v", validatorErr)
		log.Fatal(validatorErr)
	}
	logger.Info("Init user validator success", nil)

	// Create user validator
	logger.Info("Init user validator", nil)
	userValidator := userVD.NewUserValidator(validator)
	logger.Info("Init user validator success", nil)

	// Create user service
	userServive := userV1.NewService(userRepo, logger.Log, userValidator)
	if err := runServer(userServive); err != nil {
		logger.Error("server error %v", err)
		log.Fatal(err)
	}
}

func setupGrpcServerOptions() (opts []grpc.ServerOption) {
	return opts
}

func setupGrpcDialOptions() (opts []grpc.DialOption) {
	opts = append(opts, grpc.WithBlock(), grpc.WithInsecure())
	return opts
}

func setupServeMuxOptions() (opts []runtime.ServeMuxOption) {
	return opts
}

// grpcHandlerFunc returns an http.Handler that delegates to grpcServer on incoming gRPC
// connections or otherHandler otherwise. Copied from cockroachdb.
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

// https://mycodesmells.com/post/grpc-gateway-error-handler
func CustomHTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {

	const fallback = `{"error": "failed to marshal error message"}`

	st := status.Convert(err)
	sc := runtime.HTTPStatusFromCode(status.Code(err))

	errDetails := errorBody{
		Code:    st.Proto().GetCode(),
		Message: st.Message(),
		Details: make(map[string]string),
	}

	for _, detail := range st.Details() {
		switch t := detail.(type) {
		case *errdetails.BadRequest:
			for _, violation := range t.GetFieldViolations() {
				errDetails.Details[violation.GetField()] = violation.GetDescription()
			}
		}
	}

	w.Header().Set("Content-type", marshaler.ContentType())
	w.WriteHeader(sc)

	jErr := json.NewEncoder(w).Encode(errDetails)

	if jErr != nil {
		w.Write([]byte(fallback))
	}
}

/*
https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/adding_annotations/
https://github.com/johanbrandhorst/grpc-gateway-boilerplate
https://github.com/philips/grpc-gateway-example/blob/master/cmd/serve.go#L51
*/
func runServer(service userV1.Service) error {

	// Handle custom http error return
	// runtime.HTTPError = CustomHTTPError

	// Create new context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// gRPC Server
	// Create Server listener
	lis, lisErr := net.Listen("tcp", gRPCAddr)
	if lisErr != nil {
		return lisErr
	}

	// Create error channel
	errChan := make(chan error)

	// Create grpcServer with options
	grpcServerOpts := setupGrpcServerOptions()
	grpcServer := grpc.NewServer(grpcServerOpts...)

	// Register service
	userV1PB.RegisterUserServiceServer(grpcServer, service)

	// Run gRPC server
	go func() {
		logger.Info("gRPC server is listening on: %v", gRPCAddr)
		errChan <- grpcServer.Serve(lis)
	}()

	// gRPC-Gateway
	// Create gRPC gateway
	grpcDialOpts := setupGrpcDialOptions()
	conn, gatewayErr := grpc.DialContext(
		ctx,
		gRPCAddr,
		grpcDialOpts...,
	)
	if gatewayErr != nil {
		return gatewayErr
	}

	// Create Server Mux
	gwServerMuxOpts := setupServeMuxOptions()
	gwmux := runtime.NewServeMux(gwServerMuxOpts...)

	// Register ServiceHandler
	if handlerErr := userV1PB.RegisterUserServiceHandler(ctx, gwmux, conn); handlerErr != nil {
		return handlerErr
	}

	gwServer := http.Server{
		Addr:    gwAddr,
		Handler: grpcHandlerFunc(grpcServer, gwmux),
	}

	// Run gRPC-gateway server
	go func() {
		logger.Info("Gateway listening on: %v", gwAddr)
		errChan <- gwServer.ListenAndServe()
	}()

	// Handle shutdown signal
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		logger.Info("got %v signal, graceful shutdown server", <-c)
		cancel()
		grpcServer.GracefulStop()
		gwServer.Shutdown(ctx)
		close(errChan)
	}()

	// Waiting for signal
	serErr := <-errChan
	return serErr
}
