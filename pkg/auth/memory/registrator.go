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
func (i *registerRepository) Register(newUser auth.HashedCreateUser) (*auth.User, error) {
	if newUser.HashedPassword == "" {
		return nil, auth.ErrPwdNotSet
	}

	i.newID++
	usr := userRecord{
		ID:        i.newID,
		UID:       auth.UID(ksuid.New().String()),
		Username:  newUser.Username,
		Email:     newUser.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Password:  newUser.HashedPassword,
	}

	i.store.Add(usr.ID.String(), usr)

	if value, ok := i.store.Get(usr.ID.String()); ok {
		if record, ok := value.(userRecord); ok {
			return recordToUser(record), nil
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
