package postgres

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/message"
)

type service struct {
	message.Finder
	message.Deleter
	message.Creator
}

// NewService creates a new user service
func NewService(db *sql.DB) message.Service {
	r := NewFinder(db)
	d := NewDeleter(db)
	c := NewCreator(db)

	return &service{
		Finder:  *r,
		Deleter: *d,
		Creator: *c,
	}
}
