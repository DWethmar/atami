package auth

import "github.com/dwethmar/atami/pkg/model"

// Service defines interations with users
type Service interface {
	Authenticate(credentials Credentials) (bool, error)
	FindAll() ([]*model.User, error)
	FindByID(ID model.UserID) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	Delete(ID model.UserID) error
	Register(newUser CreateUser) (*model.User, error)
	ValidateUser(user User) error
	ValidateNewUser(newUser CreateUser) error
}

type service struct {
	Authenticator
	Finder
	Deleter
	Registrator
	Validator
}

// NewService creates a new user service
func NewService(
	a Authenticator,
	f Finder,
	d Deleter,
	r Registrator,
	v Validator,
) Service {
	return &service{
		Authenticator: a,
		Finder:        f,
		Deleter:       d,
		Registrator:   r,
		Validator:     v,
	}
}
