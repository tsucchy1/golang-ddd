
package config

import (
	"interfaces/handler"
	"application/usecase"
	"domain/repository"
	"infrastructure/persistence"
)

type Registry interface {
	NewUserRepository() repository.UserRepository
	NewUserUseCase() usecase.UserUseCase
	NewAppHandler() 
}

type registry struct {}

func (r *registry) NewUserRepository() repository.UserRepository {
	return &persistence.NewUserRepository()
}

func (r *registry) NewUserUseCase() usecase.UserUseCase {
	repo := r.NewUserRepository()
	return &usecase.NewUserUseCase(repo)
}

func (r *registry) NewAppHandler() {
	return &AppHandler{
		handler.NewGetUserHandler(r.NewUserUseCase())
	}
}