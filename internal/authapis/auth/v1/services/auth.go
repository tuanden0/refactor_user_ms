package services

import authv1 "github.com/tuanden0/refactor_user_ms/proto/gen/go/authen/v1"

type Service interface {
	authv1.AuthServiceServer
}

type service struct {
	authv1.UnimplementedAuthServiceServer
}

func NewService() Service {
	return &service{}
}
