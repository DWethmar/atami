package memory

import (
	"errors"

	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/dwethmar/atami/pkg/domain/user/memory/util"
	"github.com/dwethmar/atami/pkg/memstore"
)

var layoutISO = "2006-01-02"

// creatorRepository stores new messages
type creatorRepository struct {
	store *memstore.Memstore
	newID int
}

// Create new user
func (i *creatorRepository) Create(newUser user.CreateUser) (*user.User, error) {
	userStore := i.store.GetUsers()
	users, err := userStore.All()
	if err != nil {
		return nil, err
	}

	// Check for user with same username or email
	if match, err := filterList(users, func(record user.User) bool {
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
	usr := user.User{
		ID:        i.newID,
		UID:       newUser.UID,
		Username:  newUser.Username,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
		Password:  newUser.Password,
	}

	userStore.Put(usr.ID, util.ToMemory(usr))

	if r, ok := userStore.Get(usr.ID); ok {
		user := util.FromMemory(r)
		return &user, nil
	}

	return nil, errors.New("could not register user")
}

// NewCreator creates new creator.
func NewCreator(
	store *memstore.Memstore,
	finder *user.Finder,
) *user.Creator {
	return user.NewCreator(&creatorRepository{store, 0}, finder)
}
