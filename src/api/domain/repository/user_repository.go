package repository

import (
	"api/domain/model"
)

type UserRepository interface {
	FindOne(name string) (*model.User, error)
}
