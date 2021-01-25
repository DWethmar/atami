package memory

import (
	"errors"
	"strconv"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
)

var layoutISO = "2006-01-02"

// creatorRepository stores new messages
type creatorRepository struct {
	store *memstore.Store
	newID int
}

// Create new user
func (i *creatorRepository) Create(newUser user.CreateAction) (*user.User, error) {
	if newUser.Password == "" {
		return nil, user.ErrPwdNotSet
	}

	users := i.store.GetUsers()

	// Check for user with same username or email
	if match, err := filterList(users.All(), func(record userRecord) bool {
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
		UID:       newUser.UID,
		Username:  newUser.Username,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
		Password:  newUser.Password,
	}
	IDStr := strconv.Itoa(usr.ID)
	users.Put(IDStr, usr)

	if value, ok := users.Get(IDStr); ok {
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
	return user.NewCreator(&creatorRepository{store, 0})
}
