
package repository

import (
	"context"
	"domain/model"
)

type UserRepository interface {
    FindOne(name string) (*model.User, error)
}