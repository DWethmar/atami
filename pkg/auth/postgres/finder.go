package postgres

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/auth"
)

var getUsers = `
SELECT
	id,  
	uid,
	username, 
	email,
	created_on, 
	updated_on
FROM public.user
`

var getUserByID = `
SELECT
	id,  
	uid,
	username, 
	email,
	created_on, 
	updated_on
FROM public.user
WHERE id = $1
LIMIT 1`

var getUserByEmail = `
SELECT
	id,  
	uid,
	username, 
	email,
	created_on, 
	updated_on
FROM public.user
WHERE email = $1
LIMIT 1`

var getUserByUsername = `
SELECT
	id,  
	uid,
	username, 
	email,
	created_on, 
	updated_on
FROM public.user
WHERE username = $1
LIMIT 1`

// findRepository reads messages from memory
type findRepository struct {
	db *sql.DB
}

// FindAll get multiple messages
func (f findRepository) FindAll() ([]*auth.User, error) {
	rows, err := f.db.Query(getUsers)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	entries := make([]*auth.User, 0)

	for rows.Next() {
		entry := &auth.User{}

		if err := rows.Scan(
			&entry.ID,
			&entry.UID,
			&entry.Username,
			&entry.Email,
			&entry.CreatedAt,
			&entry.UpdatedAt,
		); err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}

// FindByID get one message
func (f findRepository) FindByID(ID auth.ID) (*auth.User, error) {
	entry := &auth.User{}
	if err := f.db.QueryRow(getUserByID, ID).Scan(
		&entry.ID,
		&entry.UID,
		&entry.Username,
		&entry.Email,
		&entry.CreatedAt,
		&entry.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, auth.ErrCouldNotFind
		}
		return nil, err
	}
	return entry, nil
}

// FindByEmail func
func (f *findRepository) FindByEmail(email string) (*auth.User, error) {
	entry := &auth.User{}
	if err := f.db.QueryRow(getUserByEmail, email).Scan(
		&entry.ID,
		&entry.UID,
		&entry.Username,
		&entry.Email,
		&entry.CreatedAt,
		&entry.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, auth.ErrCouldNotFind
		}
		return nil, err
	}
	return entry, nil
}

// FindByEmail func
func (f *findRepository) FindByUsername(username string) (*auth.User, error) {
	entry := &auth.User{}
	if err := f.db.QueryRow(getUserByUsername, username).Scan(
		&entry.ID,
		&entry.UID,
		&entry.Username,
		&entry.Email,
		&entry.CreatedAt,
		&entry.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, auth.ErrCouldNotFind
		}
		return nil, err
	}
	return entry, nil
}

// NewFinder return a new in memory listin repository
func NewFinder(db *sql.DB) *auth.Finder {
	return auth.NewFinder(&findRepository{db})
}
