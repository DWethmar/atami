package postgres

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/user"
)

// findRepository reads messages from memory
type findRepository struct {
	db *sql.DB
}

// FindAll get multiple messages
func (f findRepository) Find() ([]*user.User, error) {
	return querySelectUsers(f.db, 100, 0)
}

// FindByID get one message by ID
func (f findRepository) FindByID(ID int) (*user.User, error) {
	return queryRowSelectUserByID(f.db, ID)
}

// FindByUID get one message by UID
func (f findRepository) FindByUID(UID string) (*user.User, error) {
	return queryRowSelectUserByUID(f.db, UID)
}

// FindByEmail func
func (f *findRepository) FindByEmail(email string) (*user.User, error) {
	return queryRowSelectUserByEmail(f.db, email)
}

// FindByEmailWithPassword func
func (f *findRepository) FindByEmailWithPassword(email string) (*user.User, error) {
	return queryRowSelectUserByEmailWithPassword(f.db, email)
}

// FindByEmail func
func (f *findRepository) FindByUsername(username string) (*user.User, error) {
	return queryRowSelectUserByUsername(f.db, username)
}

// NewFinder return a new in memory listin repository
func NewFinder(db *sql.DB) *user.Finder {
	return user.NewFinder(&findRepository{db})
}
