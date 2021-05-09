package postgres

import (
	"github.com/dwethmar/atami/pkg/domain/message"
)

func defaultMap(row Row) (*message.Message, error) {
	e := &message.Message{}
	if err := row.Scan(
		&e.ID,
		&e.UID,
		&e.Text,
		&e.CreatedByUserID,
		&e.CreatedAt,
		&e.UpdatedAt,
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
		&e.UpdatedAt,
		&e.User.ID,
		&e.User.UID,
		&e.User.Username,
	); err != nil {
		return nil, err
	}
	return e, nil
}
