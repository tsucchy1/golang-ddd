
package persistence

import (
	"api/domain/model"
	"api/domain/repository"
	"api/library/err"
)

type userRepository struct {}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindOne(name string) (*model.User, error) {

	users := []map[string]string{
		{ "name": "apple" },
		{ "name": "google" },
		{ "name": "tesla" },
		{ "name": "toyota" },
	}

	for _, u := range users {
		if u["name"] == name {
			return &model.User{name}, nil
		}
	}

	return nil, &err.RecordNotFound{name, "User", nil}
}