package postgres

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/domain/user"
)

func defaultMap(row Row) (*user.User, error) {
	e := &user.User{}

	var biography sql.NullString

	if err := row.Scan(
		&e.ID,
		&e.UID,
		&e.Username,
		&e.Email,
		&biography,
		&e.CreatedAt,
		&e.UpdatedAt,
	); err != nil {
		return nil, err
	}

	if biography.Valid {
		e.Biography = biography.String
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

	var biography sql.NullString

	if err := row.Scan(
		&e.ID,
		&e.UID,
		&e.Username,
		&e.Email,
		&biography,
		&e.CreatedAt,
		&e.UpdatedAt,
		&e.Password,
	); err != nil {
		return nil, err
	}

	if biography.Valid {
		e.Biography = biography.String
	}

	return e, nil
}
