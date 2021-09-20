package services

import (
	"github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/repositories"
	"github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/validators"
	userV1PB "github.com/tuanden0/refactor_user_ms/proto/gen/go/user/v1"
	"go.uber.org/zap"
)

type Service interface {
	userV1PB.UserServiceServer
}

type service struct {
	userV1PB.UnimplementedUserServiceServer
	repo      repositories.UserRepository
	log       *zap.Logger
	validator *validators.UserValidator
}

func NewService(repo repositories.UserRepository, log *zap.Logger, vd *validators.UserValidator) Service {
	return &service{
		repo:      repo,
		log:       log,
		validator: vd,
	}
}
