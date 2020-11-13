package message

import "github.com/dwethmar/atami/pkg/model"

// Service defines interations with users
type Service interface {
	FindByID(ID model.MessageID) (*model.Message, error)
	FindAll() ([]*model.Message, error)
	Delete(ID model.MessageID) error
	Create(newMessage NewMessage) (*model.Message, error)
}

type service struct {
	Finder
	Deleter
	Creator
}

// NewService creates a new user service
func NewService(
	r Finder,
	d Deleter,
	c Creator,
) Service {
	return &service{
		Finder:  r,
		Deleter: d,
		Creator: c,
	}
}
