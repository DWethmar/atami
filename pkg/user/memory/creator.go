package memory

import (
	"errors"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
	"github.com/segmentio/ksuid"
)

var layoutISO = "2006-01-02"

// creatorRepository stores new messages
type creatorRepository struct {
	store *memstore.Store
	newID user.ID
}

func checkUniqueEmail(email string, users []*user.User) bool {
	for _, user := range users {
		if user.Email == email {
			return false
		}
	}
	return true
}

// Create new user
func (i creatorRepository) Create(newUser user.NewUser) (*user.User, error) {
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
			return nil, errors.New("Error while parsing user")
		}
	}

	if unique := checkUniqueEmail(newUser.Email, users); !unique {
		return nil, user.ErrEmailAlreadyTaken
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
		return nil, errors.New("Error parsing user")
	}

	return nil, errors.New("Error while finding user in memory")
}

// NewCreator creates new messages.
func NewCreator(validator *user.Validator, store *memstore.Store) *user.Creator {
	return user.NewCreator(
		&creatorRepository{
			store,
			0,
		},
		validator,
	)
}
