
package usecase

import (
	"application/input"
	"application/output"
	"domain/repository"
)

type UserUseCase interface {
	GetUser(name string) (*model.User, error)
}

type userUseCase struct {
	repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUseCase{r}
}

func (u *userUseCase) GetUser(i input.GetUser) (output.User, error) {
	user, err := u.FindOne(i.Name)
	if err != nil {
		return nil, err
	}

	return output.User{user.Name}, nil
}