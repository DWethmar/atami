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
	rows, err := f.db.Query(selectUsers)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	entries := make([]*user.User, 0)

	for rows.Next() {
		entry := &user.User{}

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

// FindByID get one message by ID
func (f findRepository) FindByID(ID int) (*user.User, error) {
	entry := &user.User{}
	if err := f.db.QueryRow(selectUserByID, ID).Scan(
		&entry.ID,
		&entry.UID,
		&entry.Username,
		&entry.Email,
		&entry.CreatedAt,
		&entry.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, user.ErrCouldNotFind
		}
		return nil, err
	}
	return entry, nil
}

// FindByUID get one message by UID
func (f findRepository) FindByUID(UID string) (*user.User, error) {
	entry := &user.User{}
	if err := f.db.QueryRow(selectUserByUID, UID).Scan(
		&entry.ID,
		&entry.UID,
		&entry.Username,
		&entry.Email,
		&entry.CreatedAt,
		&entry.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, user.ErrCouldNotFind
		}
		return nil, err
	}
	return entry, nil
}

// FindByEmail func
func (f *findRepository) FindByEmail(email string) (*user.User, error) {
	entry := &user.User{}
	if err := f.db.QueryRow(selectUserByEmail, email).Scan(
		&entry.ID,
		&entry.UID,
		&entry.Username,
		&entry.Email,
		&entry.CreatedAt,
		&entry.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, user.ErrCouldNotFind
		}
		return nil, err
	}
	return entry, nil
}

// FindByEmailWithPassword func
func (f *findRepository) FindByEmailWithPassword(email string) (*user.User, error) {
	entry := &user.User{}
	if err := f.db.QueryRow(selectUserByEmailWithPassword, email).Scan(
		&entry.ID,
		&entry.UID,
		&entry.Username,
		&entry.Email,
		&entry.CreatedAt,
		&entry.UpdatedAt,
		&entry.Password,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, user.ErrCouldNotFind
		}
		return nil, err
	}
	return entry, nil
}

// FindByEmail func
func (f *findRepository) FindByUsername(username string) (*user.User, error) {
	entry := &user.User{}
	if err := f.db.QueryRow(selectUserByUsername, username).Scan(
		&entry.ID,
		&entry.UID,
		&entry.Username,
		&entry.Email,
		&entry.CreatedAt,
		&entry.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, user.ErrCouldNotFind
		}
		return nil, err
	}
	return entry, nil
}

// NewFinder return a new in memory listin repository
func NewFinder(db *sql.DB) *user.Finder {
	return user.NewFinder(&findRepository{db})
}
