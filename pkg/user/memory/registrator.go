package memory

import (
	"errors"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
	"github.com/segmentio/ksuid"
)

var layoutISO = "2006-01-02"

// registerRepository stores new messages
type registerRepository struct {
	store *memstore.Store
	newID user.ID
}

// Create new user
func (i registerRepository) Register(newUser user.NewUser) (*user.User, error) {
	if newUser.Password == "" {
		return nil, user.ErrPwdNotSet
	}

	// Check if unique email
	results := i.store.List()
	users := make([]*user.User, len(results))
	for i, l := range results {
		if item, ok := l.(user.User); ok {
			users[i] = &item
		} else {
			return nil, errCouldNotParse
		}
	}

	i.newID++
	usr := user.User{
		ID:        i.newID,
		UID:       user.UID(ksuid.New().String()),
		Username:  newUser.Username,
		Email:     newUser.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Password:  newUser.Password,
	}
	i.store.Add(string(usr.UID), usr)

	if value, ok := i.store.Get(string(usr.UID)); ok {
		if usrResult, ok := value.(user.User); ok {
			return &usrResult, nil
		}
		return nil, errCouldNotParse
	}

	return nil, errors.New("error while finding user")
}

// NewRegistrator creates new registartor.
func NewRegistrator(
	finder *user.Finder,
	validator *user.Validator,
	store *memstore.Store,
) *user.Registrator {
	return user.NewRegistartor(
		&registerRepository{
			store,
			0,
		},
		finder,
		validator,
	)
}
