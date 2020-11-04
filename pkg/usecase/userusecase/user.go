package userusecase

import (
	"errors"

	"github.com/dwethmar/atami/pkg/model"
	"github.com/dwethmar/atami/pkg/user"
)

// Usecase struct defines interactions with users
type Usecase struct {
	service user.Service
}

// NewUserUsecase creates a new NewUserUsecase
func NewUserUsecase(service user.Service) *Usecase {
	return &Usecase{
		service: service,
	}
}

// List lists users
func (u *Usecase) List() ([]*model.User, error) {
	users, err := u.service.FindAll()
	if err != nil {
		return nil, err
	}
	return toUsers(users), nil
}

// Register registers user
func (u *Usecase) Register(username string, email string, password string) (*model.User, error) {
	if user, err := u.service.FindByEmail(email); user != nil || err != nil {
		if err != nil {
			return nil, err
		}

		if user != nil {
			return nil, errors.New("email already taken")
		}
	}

	user, err := u.service.Create(user.NewUser{
		Username: username,
		Email:    email,
		Password: password,
	})

	if err == nil {
		return toUser(user), nil
	}

	return nil, err
}
