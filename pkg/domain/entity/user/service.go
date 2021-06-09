package user

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

//Create a message
func (s *Service) Create(e Create) (entity.ID, error) {
	return s.repo.Create(&User{
		UID:             entity.NewUID(),
		CreatedAt:       time.Now(),
	})
}

//Get a message
func (s *Service) Get(id entity.ID) (*User, error) {
	return s.repo.Get(id)
}

//List messages
func (s *Service) List(limit, offset uint) ([]*User, error) {
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
func (s *Service) Update(ID entity.ID, e Update) error {
	user, err := s.Get(ID)
	if err != nil {
		return err
	}
	user.Biography = e.Biography
	user.UpdatedAt = time.Now()
	return s.repo.Update(user)
}
