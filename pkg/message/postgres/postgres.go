package postgres

import (
	"errors"

	"github.com/dwethmar/atami/pkg/message"
)

//go:generate go run ./generate/gen.go

var (
	errCouldNotParse = errors.New("could not parse user")
)

func defaultMap(row Row) (*message.Message, error) {
	e := &message.Message{}
	if err := row.Scan(
		&e.ID,
		&e.UID,
		&e.Text,
		&e.CreatedByUserID,
		&e.CreatedAt,
	); err != nil {
		return nil, err
	}
	return e, nil
}

func mapMessageWithUser(row Row) (*message.Message, error) {
	e := &message.Message{
		User: &message.User{},
	}
	if err := row.Scan(
		&e.ID,
		&e.UID,
		&e.Text,
		&e.CreatedByUserID,
		&e.CreatedAt,
		&e.User.ID,
		&e.User.UID,
		&e.User.Username,
	); err != nil {
		return nil, err
	}
	return e, nil
}
