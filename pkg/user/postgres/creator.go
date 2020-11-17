package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/dwethmar/atami/pkg/user"
	"github.com/lib/pq"

	"github.com/segmentio/ksuid"
)

var checkUniqueUsername = fmt.Sprintf(`
SELECT 1
FROM %s 
WHERE username = $1 
LIMIT 1`, tableName)

var checkUniqueEmail = fmt.Sprintf(`
SELECT 1
FROM %s 
WHERE email = $1
LIMIT 1`, tableName)

var insertUser = fmt.Sprintf(`
INSERT INTO %s (
	uid,
	username, 
	email,
	password,
	created_at, 
	updated_at
)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`, tableName)

// createRepository stores new messages
type creatorRepository struct {
	db *sql.DB
}

func isUniqueUsername(db *sql.DB, username string) (bool, error) {
	stmt, err := db.Prepare(checkUniqueUsername)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var result int
	if err = stmt.QueryRow(
		username,
	).Scan(&result); err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
	}

	return result == 0, nil
}

func isUniqueEmail(db *sql.DB, email string) (bool, error) {
	stmt, err := db.Prepare(checkUniqueEmail)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var result int
	if err = stmt.QueryRow(
		email,
	).Scan(&result); err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
	}

	return result == 0, nil
}

// Create new user
func (i *creatorRepository) Create(newUser user.CreateUser) (*user.User, error) {
	if newUser.HashedPassword == "" {
		return nil, user.ErrPwdNotSet
	}

	if unique, err := isUniqueUsername(i.db, newUser.Username); err == nil {
		if !unique {
			return nil, user.ErrUsernameAlreadyTaken
		}
	} else if err != nil {
		return nil, err
	}

	if unique, err := isUniqueEmail(i.db, newUser.Email); err == nil {
		if !unique {
			return nil, user.ErrEmailAlreadyTaken
		}
	} else if err != nil {
		return nil, err
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
		if err, ok := err.(*pq.Error); ok {
			fmt.Println("pq error:", err.Code.Name())
		}
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
	db *sql.DB,
) *user.Creator {
	return user.NewCreator(
		&creatorRepository{
			db,
		},
		user.NewDefaultValidator(),
	)
}
