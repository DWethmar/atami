package message

import (
	"errors"

	"github.com/dwethmar/atami/pkg/model"
)

var (
	// ErrCouldNotFind error
	ErrCouldNotFind = errors.New("could not find message")
)

// FindRepository defines a messsage listing repository
type FindRepository interface {
	FindByID(ID model.MessageID) (*Message, error)
	Find() ([]*Message, error)
}

// Finder lists messages.
type Finder struct {
	readerRepo FindRepository
}

// FindByID return a list of list items.
func (m *Finder) FindByID(ID model.MessageID) (*model.Message, error) {
	message, err := m.readerRepo.FindByID(ID)
	if err != nil {
		return nil, err
	}
	return ToMessage(message), nil
}

// Find return a list of list items.
func (m *Finder) Find() ([]*model.Message, error) {
	results, err := m.readerRepo.Find()
	if err != nil {
		return nil, err
	}

	messages := make([]*model.Message, len(results))
	for i, result := range results {
		messages[i] = ToMessage(result)
	}

	return messages, nil
}

// NewFinder returns a new Listing
func NewFinder(r FindRepository) *Finder {
	return &Finder{r}
}
