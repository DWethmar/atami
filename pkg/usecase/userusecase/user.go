package userusecase

import (
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
	_, err := u.service.FindByEmail(email)
	if err != nil {
		if err != user.ErrCouldNotFind {
			return nil, err
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
