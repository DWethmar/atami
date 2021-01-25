package postgres

import (
	"github.com/dwethmar/atami/pkg/user"
)

func defaultMap(row Row) (*user.User, error) {
	e := &user.User{}
	if err := row.Scan(
		&e.ID,
		&e.UID,
		&e.Username,
		&e.Email,
		&e.Biography,
		&e.CreatedAt,
		&e.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return e, nil
}

func mapIsUniqueCheck(row Row) (bool, error) {
	i := 0
	row.Scan(&i)
	return i == 0, row.Err()
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
