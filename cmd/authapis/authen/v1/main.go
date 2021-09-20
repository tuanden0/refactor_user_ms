package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"github.com/tuanden0/refactor_user_ms/internal/authapis/auth/v1/services"
	logger "github.com/tuanden0/refactor_user_ms/internal/logs/logrus_driver"
	authv1 "github.com/tuanden0/refactor_user_ms/proto/gen/go/authen/v1"
	"google.golang.org/grpc"
)

const version = "1.0.0"

type authenApplication struct {
	network        string
	grpcAddr       string
	gwAddr         string
	log            *logrus.Logger
	version        string
	service        authv1.AuthServiceServer
	grpcServerOpts []grpc.ServerOption
	grpcDialOpts   []grpc.DialOption
	serverpMuxOpts []runtime.ServeMuxOption
}

func main() {

	app := authenApplication{}
	app.version = version
	app.network = "tcp"

	// Get addr
	flag.StringVar(&app.grpcAddr, "grpc", ":50002", "gRPC server address")
	flag.StringVar(&app.gwAddr, "gw", ":8001", "HTTP server address")
	flag.Parse()

	// Get log
	app.log = logger.Log
	logFile, logErr := os.OpenFile("logs/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if logErr != nil {
		log.Fatalf("failed to create log %v", logErr)
	}
	defer logFile.Close()
	app.log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	// Get service
	app.service = services.NewService()

	// Get opts
	app.grpcServerOpts = setupGrpcServerOptions()
	app.grpcDialOpts = setupGrpcDialOptions()
	app.serverpMuxOpts = setupServeMuxOptions()

	if err := app.runServer(); err != nil {
		logger.Error("server error %v", err)
		log.Fatal(err)
	}
}

func (a *authenApplication) runServer() error {

	// Create new context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	lis, lisErr := net.Listen(a.network, a.grpcAddr)
	if lisErr != nil {
		return lisErr
	}

	errChan := make(chan error)

	grpcServer := grpc.NewServer(a.grpcServerOpts...)
	authv1.RegisterAuthServiceServer(grpcServer, a.service)

	go func() {
		logger.Info(fmt.Sprintf("gRPC server is listening on: %v", a.grpcAddr))
		errChan <- grpcServer.Serve(lis)
	}()

	conn, gatewayErr := grpc.DialContext(
		ctx,
		a.grpcAddr,
		a.grpcDialOpts...,
	)
	if gatewayErr != nil {
		return gatewayErr
	}

	gwmux := runtime.NewServeMux(a.serverpMuxOpts...)
	if err := authv1.RegisterAuthServiceHandler(ctx, gwmux, conn); err != nil {
		return err
	}

	gwServer := http.Server{
		Addr:    a.gwAddr,
		Handler: grpcHandlerFunc(grpcServer, gwmux),
	}
	go func() {
		logger.Info(fmt.Sprintf("Gateway listening on: %v", a.gwAddr))
		errChan <- gwServer.ListenAndServe()
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		logger.Info(fmt.Sprintf("got %v signal, graceful shutdown server", <-c))
		cancel()
		grpcServer.GracefulStop()
		gwServer.Shutdown(ctx)
		close(errChan)
	}()

	serErr := <-errChan
	return serErr
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

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}
