package postgres

import (
	"errors"

	"github.com/dwethmar/atami/pkg/node"
)

var (
	errCouldNotParse = errors.New("could not parse user")
)

func defaultMap(row Row) (*node.Node, error) {
	e := &node.Node{}
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

func mapNodeWithUser(row Row) (*node.Node, error) {
	e := &node.Node{
		User: &node.User{},
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
