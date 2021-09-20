package helpers

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	MappingError  = status.Error(codes.Internal, "failed to process your data")
	CreateError   = status.Error(codes.Internal, "failed to create user")
	RetrieveError = status.Error(codes.Internal, "failed to retrieve user")
	UpdateError   = status.Error(codes.Internal, "failed to update user")
	DeleteError   = status.Error(codes.Internal, "failed to delete user")
	ListError     = status.Error(codes.Internal, "failed to retrieve users")
)
