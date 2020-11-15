package memory

import (
	"errors"
	"strconv"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
	"github.com/segmentio/ksuid"
)

var layoutISO = "2006-01-02"

// creatorRepository stores new messages
type creatorRepository struct {
	store *memstore.Store
	newID int
}

// Create new user
func (i *creatorRepository) Create(newUser user.CreateUser) (*user.User, error) {
	if newUser.HashedPassword == "" {
		return nil, user.ErrPwdNotSet
	}

	// Check for user with same username or email
	if match, err := filterList(i.store.List(), func(record userRecord) bool {
		return newUser.Username == record.Username || newUser.Email == record.Email
	}); match != nil {

		if match.Email == newUser.Email {
			return nil, user.ErrEmailAlreadyTaken
		}

		if match.Username == newUser.Username {
			return nil, user.ErrUsernameAlreadyTaken
		}

		return nil, err
	} else if err != nil && err != user.ErrCouldNotFind {
		return nil, err
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

// NewCreator creates new creator.
func NewCreator(
	store *memstore.Store,
) *user.Creator {
	return user.NewCreator(
		&creatorRepository{
			store,
			0,
		},
		user.NewDefaultValidator(),
	)
}
