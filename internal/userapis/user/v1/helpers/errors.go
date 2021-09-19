package helpers

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	MappingError = status.Error(codes.Internal, "failed to process your data")
	CreateError  = status.Error(codes.Internal, "failed to create user")
)
