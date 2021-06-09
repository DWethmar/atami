package user

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/dwethmar/atami/pkg/domain"
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

// Authenticated result
type Authenticated struct {
	AccessToken string
	RefreshToken string
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

	if usr, err := s.repo.GetByEmail(e.Email); usr != nil && err == nil {
		return 0, ErrEmailAlreadyTaken
	} else if err != nil && err != domain.ErrNotFound{
		return 0, err
	}

	if usr, err := s.repo.GetByUsername(e.Username); usr != nil && err == nil {
		return 0, ErrUsernameAlreadyTaken
	} else if err != nil && err != domain.ErrNotFound {
		return 0, err
	}

	return s.repo.Create(&User{
		UID:		entity.NewUID(),
		Username: 	e.Username,
		Email: 		e.Email,
		Biography:  e.Biography,
		Password: 	HashPassword([]byte(e.Password)),
		CreatedAt:	entity.Now(),
		UpdatedAt: 	entity.Now(),
	})
}

//Get a message
func (s *Service) Get(ID entity.ID) (*User, error) {
	return s.repo.Get(ID)
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
func (s *Service) Update(ID entity.ID, e *Update) error {
	user, err := s.Get(ID)
	if err != nil {
		return err
	}
	user.Biography = e.Biography
	user.UpdatedAt = time.Now()
	return s.repo.Update(user)
}

func (s *Service) Authenticate(email, password string) (*Authenticated, error) {
	if email == "" {
		return nil, ErrEmailRequired
	}

	if password == "" {
		return nil, ErrPasswordRequired
	}

	credentials, err := s.repo.GetCredentials(email)
	if err != nil {
		if err != domain.ErrNotFound {
			return nil, err
		}
		return nil, errors.New("could not authenticate")
	}

	if !ComparePasswords(password, []byte(credentials.Password)) {
		return nil, errors.New("could not authenticate")
	}

	session := strconv.FormatInt(time.Now().UnixNano(), 10)

	accessTokenDuration := time.Minute * 60
	accessToken, err := CreateAccessToken(credentials.UID, session, time.Now().Add(accessTokenDuration).Unix())
	if err != nil {
		return nil, err
	}

	refreshTokenDuration := time.Hour * 730
	refreshToken, err := CreateRefreshToken(credentials.UID, session, time.Now().Add(refreshTokenDuration).Unix())
	if err != nil {
		return nil, err
	}

	return &Authenticated{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Service) ValidateCreate(c *Create) error {
	err := errValidate{}

	if e := ValidateUsername(c.Username); e != nil {
		err.Errors = append(err.Errors, e)
	}

	if e := ValidateEmail(c.Email); e != nil {
		err.Errors = append(err.Errors, e)
	}

	if e := ValidatePassword(c.Password); e != nil {
		err.Errors = append(err.Errors, e)
	}

	if e := ValidateBiography(c.Biography); e != nil {
		err.Errors = append(err.Errors, e)
	}

	if err.Valid() {
		return nil
	}

	return err
}

func (s *Service) ValidateUpdate(u *Update) error {
	err := errValidate{}

	if e := ValidateBiography(u.Biography); e != nil {
		err.Errors = append(err.Errors, e)
	}

	if err.Valid() {
		return nil
	}

	return err
}