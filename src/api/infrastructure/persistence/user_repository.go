
package persistence

import (
	"domain/repository"
	"library/err"
)

type userRepository struct {}

func NewUserRepository() *repository.UserRepository {
	return &UserRepository{}
}

func (r *userRepository) FindOne(name string) (*model.User, error) {

	users := []map[string]string{
		{ "name": "apple" },
		{ "name": "google" },
		{ "name": "tesla" },
		{ "name": "toyota" },
	}

	for u := range users {
		if u["name"] == name {
			return &model.User{name}, nil
		}
	}

	return nil, &err.RecordNotFound{name, "User", nil}
}