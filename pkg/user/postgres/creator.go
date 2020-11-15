package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/dwethmar/atami/pkg/user"

	"github.com/segmentio/ksuid"
)

var insertUser = fmt.Sprintf(`
INSERT INTO %s (
	uid,
	username, 
	email,
	password,
	created_on, 
	updated_on
)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`, tableName)

// createRepository stores new messages
type creatorRepository struct {
	db *sql.DB
}

// Create new user
func (i *creatorRepository) Create(newUser user.CreateUser) (*user.User, error) {
	if newUser.HashedPassword == "" {
		return nil, user.ErrPwdNotSet
	}

	uid := ksuid.New().String()

	stmt, err := i.db.Prepare(insertUser)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	now := time.Now().UTC()

	var userID int
	if err = stmt.QueryRow(
		uid,
		newUser.Username,
		newUser.Email,
		newUser.HashedPassword,
		now,
		now,
	).Scan(&userID); err != nil {
		return nil, err
	}

	if userID != 0 {
		entry := &user.User{}
		if err := i.db.QueryRow(getUserByID, userID).Scan(
			&entry.ID,
			&entry.UID,
			&entry.Username,
			&entry.Email,
			&entry.CreatedAt,
			&entry.UpdatedAt,
		); err != nil {
			return nil, err
		}
		return entry, nil
	}

	return nil, errors.New("could not register user")
}

// NewCreator creates new creator.
func NewCreator(
	validator *user.Validator,
	db *sql.DB,
) *user.Creator {
	return user.NewCreator(
		&creatorRepository{
			db,
		},
		validator,
	)
}
