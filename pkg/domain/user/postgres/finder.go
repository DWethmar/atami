package postgres

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/domain/user"
)

// findRepository reads messages from memory
type findRepository struct {
	db database.Transaction
}

// FindAll get multiple messages
func (f findRepository) Find() ([]*user.User, error) {
	return querySelectUsers(f.db, 200, 0)
}

// FindByID get one message by ID
func (f findRepository) FindByID(ID int) (*user.User, error) {
	u, err := queryRowSelectUserByID(f.db, ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, user.ErrCouldNotFind
		}
		return nil, err
	}
	return u, nil
}

// FindByUID get one message by UID
func (f findRepository) FindByUID(UID string) (*user.User, error) {
	u, err := queryRowSelectUserByUID(f.db, UID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, user.ErrCouldNotFind
		}
		return nil, err
	}
	return u, nil
}

// FindByEmail func
func (f *findRepository) FindByEmail(email string) (*user.User, error) {
	u, err := queryRowSelectUserByEmail(f.db, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, user.ErrCouldNotFind
		}
		return nil, err
	}
	return u, nil
}

// FindByEmailWithPassword func
func (f *findRepository) FindByEmailWithPassword(email string) (*user.User, error) {
	u, err := queryRowSelectUserByEmailWithPassword(f.db, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, user.ErrCouldNotFind
		}
		return nil, err
	}
	return u, nil
}

// FindByEmail func
func (f *findRepository) FindByUsername(username string) (*user.User, error) {
	u, err := queryRowSelectUserByUsername(f.db, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, user.ErrCouldNotFind
		}
		return nil, err
	}
	return u, nil
}

// NewFinder return a new in memory listin repository
func NewFinder(db database.Transaction) *user.Finder {
	return user.NewFinder(&findRepository{db})
}
