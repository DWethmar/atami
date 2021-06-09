package message

import (
	"strings"
	"time"

	"github.com/dwethmar/atami/pkg/domain/entity"
)

type errValidate struct {
	Errors []error
}

func (err errValidate) Valid() bool {
	return len(err.Errors) == 0
}

func (err errValidate) Error() string {
	errors := make([]string, len(err.Errors))
	for i, e := range err.Errors {
		errors[i] = e.Error()
	}
	return strings.Join(errors, ". ")
}

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

//Create a message
func (s *Service) Create(e *Create) (entity.ID, error) {
	if err := s.ValidateCreate(e); err != nil {
		return 0, err
	}
	
	return s.repo.Create(&Message{
		UID:             entity.NewUID(),
		Text:            e.Text,
		CreatedByUserID: e.CreatedByUserID,
		CreatedAt:       entity.Now(),
		UpdatedAt: 	     entity.Now(),
	})
}

//Get a message
func (s *Service) Get(id entity.ID) (*Message, error) {
	return s.repo.Get(id)
}

//List messages
func (s *Service) List(limit, offset uint) ([]*Message, error) {
	return s.repo.List(limit, offset)
}

//Delete a message
func (s *Service) Delete(id entity.ID) error {
	_, err := s.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//Update a message
func (s *Service) Update(ID entity.ID, e *Update) error {
	if err := s.ValidateUpdate(e); err != nil {
		return err
	}

	message, err := s.Get(ID)
	if err != nil {
		return err
	}
	message.Text = e.Text
	message.UpdatedAt = time.Now()
	return s.repo.Update(message)
}

func (s *Service) ValidateCreate(c *Create) error {
	err := errValidate{}

	if e := ValidateText(c.Text); e != nil {
		err.Errors = append(err.Errors, e)
	}

	if err.Valid() {
		return nil
	}

	return err
}

func (s *Service) ValidateUpdate(u *Update) error {
	err := errValidate{}

	if e := ValidateText(u.Text); e != nil {
		err.Errors = append(err.Errors, e)
	}

	if err.Valid() {
		return nil
	}

	return err
}

