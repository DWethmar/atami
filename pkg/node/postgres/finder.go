package postgres

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/node"
)

// findRepository reads nodes from memory
type findRepository struct {
	db *sql.DB
}

// FindAll get multiple nodes
func (i *findRepository) Find(limit, offset int) ([]*node.Node, error) {
	return querySelectNodes(i.db, limit, offset)
}

// FindByID get one node
func (i *findRepository) FindByUID(UID string) (*node.Node, error) {
	m, err := queryRowSelectNodeByUID(i.db, UID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, node.ErrCouldNotFind
		}
		return nil, err
	}
	return m, nil
}

// FindByID get one node
func (i *findRepository) FindByID(ID int) (*node.Node, error) {
	m, err := queryRowSelectNodeByID(i.db, ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, node.ErrCouldNotFind
		}
		return nil, err
	}
	return m, nil
}

// NewFinder return a new in memory listin reader
func NewFinder(db *sql.DB) *node.Finder {
	return node.NewFinder(&findRepository{db})
}
