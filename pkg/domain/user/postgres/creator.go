package postgres

import (
	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/domain/user"
)

// createRepository stores new messages
type creatorRepository struct {
	db database.Transaction
}

// Create new user
func (i *creatorRepository) Create(newUser user.CreateUser) (*user.User, error) {
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
	db database.Transaction,
	finder *user.Finder,
) *user.Creator {
	return user.NewCreator(&creatorRepository{db}, finder)
}
