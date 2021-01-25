package postgres

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/node"
)

// creatorRepository stores new nodes
type creatorRepository struct {
	db *sql.DB
}

// Create new node
func (i creatorRepository) Create(newMsg node.CreateAction) (*node.Node, error) {
	msg, err := queryRowInsertNode(
		i.db,
		newMsg.UID,
		newMsg.Text,
		newMsg.CreatedByUserID,
		newMsg.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

// NewCreator creates new nodes creator.
func NewCreator(db *sql.DB) *node.Creator {
	return node.NewCreator(&creatorRepository{
		db,
	})
}
