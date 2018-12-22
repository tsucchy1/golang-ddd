package config

import (
	"api/application/usecase"
	"api/domain/repository"
	"api/infrastructure/persistence"
	"api/interfaces/handler"
)

type Registry interface {
	NewUserRepository() repository.UserRepository
	NewUserUseCase() usecase.UserUseCase
	NewAppHandler() handler.AppHandler
}

type registry struct{}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewUserRepository() repository.UserRepository {
	return persistence.NewUserRepository()
}

func (r *registry) NewUserUseCase() usecase.UserUseCase {
	repo := r.NewUserRepository()
	return usecase.NewUserUseCase(repo)
}

func (r *registry) NewAppHandler() handler.AppHandler {
	return handler.NewAppHandler(
		handler.NewGetUserHandler(r.NewUserUseCase()),
	)
}
