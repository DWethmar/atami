package memory

import (
	"errors"
	"time"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/segmentio/ksuid"
)

var layoutISO = "2006-01-02"

// registerRepository stores new messages
type registerRepository struct {
	store *memstore.Store
	newID auth.ID
}

// Create new user
func (i *registerRepository) Register(newUser auth.CreateUser) (*auth.User, error) {
	if newUser.HashedPassword == "" {
		return nil, auth.ErrPwdNotSet
	}

	// Check if unique email
	results := i.store.List()
	users := make([]*auth.User, len(results))
	for i, l := range results {
		if item, ok := l.(auth.User); ok {
			users[i] = &item
		} else {
			return nil, errCouldNotParse
		}
	}

	i.newID++
	usr := storedUser{
		ID:        i.newID,
		UID:       auth.UID(ksuid.New().String()),
		Username:  newUser.Username,
		Email:     newUser.Email,
		Salt:      newUser.Salt,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Password:  newUser.HashedPassword,
	}

	i.store.Add(string(usr.UID), usr)

	if value, ok := i.store.Get(string(usr.UID)); ok {
		if usrResult, ok := value.(auth.User); ok {
			return &usrResult, nil
		}
		return nil, errCouldNotParse
	}

	return nil, errors.New("error while finding user")
}

// NewRegistrator creates new registartor.
func NewRegistrator(
	finder *auth.Finder,
	validator *auth.Validator,
	store *memstore.Store,
) *auth.Registrator {
	return auth.NewRegistartor(
		&registerRepository{
			store,
			0,
		},
		finder,
		validator,
	)
}
