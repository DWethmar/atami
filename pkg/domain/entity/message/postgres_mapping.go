package message

import "github.com/dwethmar/atami/pkg/domain/entity"

func defaultMap(row Row) (*Message, error) {
	e := &Message{}
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

func insertMap(row Row) (entity.ID, error) {
	var ID entity.ID
	if err := row.Scan(
		&ID,
	); err != nil {
		return 0, err
	}
	return ID, nil
}

func mapMessageWithUser(row Row) (*Message, error) {
	e := &Message{
		CreatedBy: User{},
	}
	if err := row.Scan(
		&e.ID,
		&e.UID,
		&e.Text,
		&e.CreatedByUserID,
		&e.CreatedAt,
		&e.CreatedBy.ID,
		&e.CreatedBy.UID,
		&e.CreatedBy.Username,
	); err != nil {
		return nil, err
	}
	return e, nil
}
