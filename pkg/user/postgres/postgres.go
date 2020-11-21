package postgres

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/user"
)

//go:generate go run ./generate/gen.go

var (
	tableName = "public.user"
)

func defaultMap(row Row) (*user.User, error) {
	e := &user.User{}
	if err := row.Scan(
		&e.ID,
		&e.UID,
		&e.Username,
		&e.Email,
		&e.CreatedAt,
		&e.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return e, nil
}

func mapIsUniqueCheck(row Row) (bool, error) {
	return sql.ErrNoRows == row.Err(), row.Err()
}

func mapWithPassword(row Row) (*user.User, error) {
	e := &user.User{}
	if err := row.Scan(
		&e.ID,
		&e.UID,
		&e.Username,
		&e.Email,
		&e.CreatedAt,
		&e.UpdatedAt,
		&e.Password,
	); err != nil {
		return nil, err
	}
	return e, nil
}
