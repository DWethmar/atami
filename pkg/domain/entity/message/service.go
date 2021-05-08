package message

import (
	"time"

	"github.com/dwethmar/atami/pkg/domain/entity"
)

//Service service interface
type Service struct {
	repo Repository
}

//NewService create new use case
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//Create a book
func (s *Service) Create(e Create) (entity.ID, error) {
	e.UID = entity.NewUID()
	return s.repo.Create(&Message{
		UID:             e.UID,
		Text:            e.Text,
		CreatedByUserID: e.CreatedByUserID,
		CreatedAt:       time.Now(),
	})
}

//Get a book
func (s *Service) Get(id entity.ID) (*Message, error) {
	return s.repo.Get(id)
}

//List books
func (s *Service) List(limit, offset uint) ([]*Message, error) {
	return s.repo.List(limit, offset)
}

//Delete a book
func (s *Service) Delete(id entity.ID) error {
	_, err := s.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//Update a book
func (s *Service) Update(ID entity.ID, e Update) error {
	message, err := s.Get(ID)
	if err != nil {
		return err
	}
	message.Text = e.Text
	message.UpdatedAt = time.Now()
	return s.repo.Update(message)
}
