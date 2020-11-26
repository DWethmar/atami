package postgres

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/user"
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

	return queryRowInsertUser(
		i.db,
		newUser.UID,
		newUser.Username,
		newUser.Email,
		newUser.Password,
		newUser.CreatedAt,
		newUser.UpdatedAt,
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
