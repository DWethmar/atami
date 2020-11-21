package postgres

import (
	"database/sql"
	"time"

	"github.com/dwethmar/atami/pkg/user"

	"github.com/segmentio/ksuid"
)

// createRepository stores new messages
type creatorRepository struct {
	db *sql.DB
}

// Create new user
func (i *creatorRepository) Create(newUser user.CreateUser) (*user.User, error) {
	if newUser.Password == "" {
		return nil, user.ErrPwdNotSet
	}

	if unique, err := queryRowSelectUsernameUniqueCheck(i.db, newUser.Username); err == nil {
		if !unique {
			return nil, user.ErrUsernameAlreadyTaken
		}
	} else if err != nil {
		return nil, err
	}

	if unique, err := queryRowSelectEmailUniqueCheck(i.db, newUser.Email); err == nil {
		if !unique {
			return nil, user.ErrEmailAlreadyTaken
		}
	} else if err != nil {
		return nil, err
	}

	uid := ksuid.New().String()
	now := time.Now().UTC()

	return queryRowInsertUser(
		i.db,
		uid,
		newUser.Username,
		newUser.Email,
		newUser.Password,
		now,
		now,
	)
}

// NewCreator creates new creator.
func NewCreator(
	db *sql.DB,
) *user.Creator {
	return user.NewCreator(
		&creatorRepository{
			db,
		},
		user.NewDefaultValidator(),
	)
}
