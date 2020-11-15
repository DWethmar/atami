package memory

import (
	"errors"
	"strconv"
	"time"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/segmentio/ksuid"
)

var layoutISO = "2006-01-02"

// registerRepository stores new messages
type registerRepository struct {
	store *memstore.Store
	newID int
}

// Create new user
func (i *registerRepository) Register(newUser auth.HashedCreateUser) (*auth.User, error) {
	if newUser.HashedPassword == "" {
		return nil, auth.ErrPwdNotSet
	}

	i.newID++
	usr := userRecord{
		ID:        i.newID,
		UID:       ksuid.New().String(),
		Username:  newUser.Username,
		Email:     newUser.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Password:  newUser.HashedPassword,
	}
	IDStr := strconv.Itoa(usr.ID)
	i.store.Add(IDStr, usr)

	if value, ok := i.store.Get(IDStr); ok {
		if record, ok := value.(userRecord); ok {
			return recordToUser(record), nil
		}
		return nil, errCouldNotParse
	}

	return nil, errors.New("could not register user")
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
