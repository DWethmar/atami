// This file was generated by robots; DO NOT EDIT.
// run: 'make generate' to regenerate this file.

package postgres

import (
	"database/sql"

	"time"

	"github.com/dwethmar/atami/pkg/user"
)

// Row needs to be implemented in the the map function.
type Row interface {
	Scan(dest ...interface{}) error
	Err() error
}

// selectUsernameUniqueCheck sql query
var selectUsernameUniqueCheck = `SELECT
	1
FROM public.users
WHERE username = $1
LIMIT 1`

func mapSelectUsernameUniqueCheck(row Row) (bool, error) {
	return mapIsUniqueCheck(row)
}

func queryRowSelectUsernameUniqueCheck(
	db *sql.DB,
	username string,
) (bool, error) {
	return mapSelectUsernameUniqueCheck(db.QueryRow(
		selectUsernameUniqueCheck,
		username,
	))
}

// selectEmailUniqueCheck sql query
var selectEmailUniqueCheck = `SELECT
	1
FROM public.users
WHERE email = $1
LIMIT 1`

func mapSelectEmailUniqueCheck(row Row) (bool, error) {
	return mapIsUniqueCheck(row)
}

func queryRowSelectEmailUniqueCheck(
	db *sql.DB,
	email string,
) (bool, error) {
	return mapSelectEmailUniqueCheck(db.QueryRow(
		selectEmailUniqueCheck,
		email,
	))
}

// insertUser sql query
var insertUser = `INSERT INTO public.users
(
	uid,
	username,
	email,
	password,
	created_at,
	updated_at
)
VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6
)
RETURNING users.id, users.uid, users.username, users.email, users.created_at, users.updated_at`

func mapInsertUser(row Row) (*user.User, error) {
	return defaultMap(row)
}

func queryRowInsertUser(
	db *sql.DB,
	UID string,
	username string,
	email string,
	password string,
	createdAt time.Time,
	updateddAt time.Time,
) (*user.User, error) {
	return mapInsertUser(db.QueryRow(
		insertUser,
		UID,
		username,
		email,
		password,
		createdAt,
		updateddAt,
	))
}

// deleteUser sql query
var deleteUser = `DELETE FROM public.users
WHERE id = $1`

func execDeleteUser(
	db *sql.DB,
	ID int,
) (sql.Result, error) {
	return db.Exec(
		deleteUser,
		ID,
	)
}

// selectUsers sql query
var selectUsers = `SELECT
	users.id,
	users.uid,
	users.username,
	users.email,
	users.created_at,
	users.updated_at
FROM public.users
ORDER BY created_at ASC
LIMIT $1
OFFSET $2`

func mapSelectUsers(row Row) (*user.User, error) {
	return defaultMap(row)
}

func querySelectUsers(
	db *sql.DB,
	limit int,
	offset int,
) ([]*user.User, error) {
	rows, err := db.Query(
		selectUsers,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	entries := make([]*user.User, 0)
	for rows.Next() {
		if entry, err := mapSelectUsers(rows); err == nil {
			entries = append(entries, entry)
		} else {
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return entries, nil
}

// selectUserByID sql query
var selectUserByID = `SELECT
	users.id,
	users.uid,
	users.username,
	users.email,
	users.created_at,
	users.updated_at
FROM public.users
WHERE id = $1
LIMIT 1`

func mapSelectUserByID(row Row) (*user.User, error) {
	return defaultMap(row)
}

func queryRowSelectUserByID(
	db *sql.DB,
	ID int,
) (*user.User, error) {
	return mapSelectUserByID(db.QueryRow(
		selectUserByID,
		ID,
	))
}

// selectUserByUID sql query
var selectUserByUID = `SELECT
	users.id,
	users.uid,
	users.username,
	users.email,
	users.created_at,
	users.updated_at
FROM public.users
WHERE uid = $1
LIMIT 1`

func mapSelectUserByUID(row Row) (*user.User, error) {
	return defaultMap(row)
}

func queryRowSelectUserByUID(
	db *sql.DB,
	UID string,
) (*user.User, error) {
	return mapSelectUserByUID(db.QueryRow(
		selectUserByUID,
		UID,
	))
}

// selectUserByEmail sql query
var selectUserByEmail = `SELECT
	users.id,
	users.uid,
	users.username,
	users.email,
	users.created_at,
	users.updated_at
FROM public.users
WHERE email = $1
LIMIT 1`

func mapSelectUserByEmail(row Row) (*user.User, error) {
	return defaultMap(row)
}

func queryRowSelectUserByEmail(
	db *sql.DB,
	email string,
) (*user.User, error) {
	return mapSelectUserByEmail(db.QueryRow(
		selectUserByEmail,
		email,
	))
}

// selectUserByEmailWithPassword sql query
var selectUserByEmailWithPassword = `SELECT
	users.id,
	users.uid,
	users.username,
	users.email,
	users.created_at,
	users.updated_at,
	password
FROM public.users
WHERE email = $1
LIMIT 1`

func mapSelectUserByEmailWithPassword(row Row) (*user.User, error) {
	return mapWithPassword(row)
}

func queryRowSelectUserByEmailWithPassword(
	db *sql.DB,
	email string,
) (*user.User, error) {
	return mapSelectUserByEmailWithPassword(db.QueryRow(
		selectUserByEmailWithPassword,
		email,
	))
}

// selectUserByUsername sql query
var selectUserByUsername = `SELECT
	users.id,
	users.uid,
	users.username,
	users.email,
	users.created_at,
	users.updated_at
FROM public.users
WHERE username = $1
LIMIT 1`

func mapSelectUserByUsername(row Row) (*user.User, error) {
	return defaultMap(row)
}

func queryRowSelectUserByUsername(
	db *sql.DB,
	username string,
) (*user.User, error) {
	return mapSelectUserByUsername(db.QueryRow(
		selectUserByUsername,
		username,
	))
}