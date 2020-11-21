package postgres

import (
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

func mapUniqueCheck(row Row) (bool, error) {
	e := 0
	if err := row.Scan(
		&e,
	); err != nil {
		return false, err
	}
	return e != 0, nil
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
