package memory

import (
	"errors"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/node"
	"github.com/dwethmar/atami/pkg/node/memory/util"
)

// creatorRepository stores new nodes
type creatorRepository struct {
	store *memstore.Store
	newID int
}

// Create new node
func (i *creatorRepository) Create(newMsg node.CreateAction) (*node.Node, error) {
	nodes := i.store.GetNodes()
	users := i.store.GetUsers()

	i.newID++
	msg := node.Node{
		ID:              i.newID,
		UID:             newMsg.UID,
		Text:            newMsg.Text,
		CreatedByUserID: newMsg.CreatedByUserID,
		CreatedAt:       newMsg.CreatedAt,
	}

	if _, ok := users.Get(msg.CreatedByUserID); !ok {
		return nil, errors.New("user not found")
	}

	nodes.Put(msg.ID, util.ToMemory(msg))
	if r, ok := nodes.Get(msg.ID); ok {
		msg := util.FromMemory(r)
		return &msg, nil
	}

	return nil, errors.New("Could not find node")
}

// NewCreator creates new nodes creator.
func NewCreator(store *memstore.Store) *node.Creator {
	return node.NewCreator(&creatorRepository{
		store,
		0,
	})
}
