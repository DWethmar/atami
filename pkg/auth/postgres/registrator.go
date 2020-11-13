package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/model"
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

// registerRepository stores new messages
type registerRepository struct {
	db *sql.DB
}

// Create new user
func (i *registerRepository) Register(newUser auth.HashedCreateUser) (*auth.User, error) {
	if newUser.HashedPassword == "" {
		return nil, auth.ErrPwdNotSet
	}

	uid := model.UserUID(ksuid.New().String())

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
		entry := &auth.User{}
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

// NewRegistrator creates new registartor.
func NewRegistrator(
	finder *auth.Finder,
	validator *auth.Validator,
	db *sql.DB,
) *auth.Registrator {
	return auth.NewRegistartor(
		&registerRepository{
			db,
		},
		finder,
		validator,
	)
}
