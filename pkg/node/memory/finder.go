package memory

import (
	"fmt"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/node"
	"github.com/dwethmar/atami/pkg/node/memory/util"
)

// readerRepository reads nodes from memory
type findRepository struct {
	store *memstore.Store
}

// FindByID get one node
func (i *findRepository) FindByUID(UID string) (*node.Node, error) {
	msg, err := filterList(i.store.GetNodes().All(), func(record node.Node) bool {
		return UID == record.UID
	})

	if msg == nil {
		return nil, err
	}

	users := i.store.GetUsers()

	if err == nil {
		if user, err := util.FindUser(users, msg.CreatedByUserID); err == nil {
			msg.User = user
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}

	return msg, nil
}

// FindByID get one node
func (i *findRepository) FindByID(ID int) (*node.Node, error) {
	nodes := i.store.GetNodes()
	users := i.store.GetUsers()

	if r, ok := nodes.Get(ID); ok {
		msg := util.FromMemory(r)
		if user, err := util.FindUser(users, msg.CreatedByUserID); err == nil {
			msg.User = user
		} else {
			return nil, err
		}

		return &msg, nil
	}
	return nil, node.ErrCouldNotFind
}

// FindAll get multiple nodes
func (i *findRepository) Find(limit, offset int) ([]*node.Node, error) {
	nodes := i.store.GetNodes()
	users := i.store.GetUsers()

	if len := nodes.Len(); len == 0 {
		return nil, nil
	} else if offset+limit > len {
		limit = len - offset
	}

	paged := nodes.Slice(offset, limit)
	items := make([]*node.Node, len(paged))

	for i, r := range paged {

		msg := util.FromMemory(r)

		fmt.Println("------------------------------------------------------------------------")
		fmt.Println(fmt.Sprintf("Created BY %d", msg.CreatedByUserID))
		fmt.Println(util.FindUser(users, msg.CreatedByUserID))
		test, _ := users.Get(msg.CreatedByUserID)
		fmt.Println(test)
		fmt.Println("------------------------------------------------------------------------")

		if user, err := util.FindUser(users, msg.CreatedByUserID); err == nil {
			msg.User = user
		} else {
			return nil, err
		}
		items[i] = &msg
	}

	return items, nil
}

// NewFinder return a new in memory listin reader
func NewFinder(store *memstore.Store) *node.Finder {
	return node.NewFinder(&findRepository{store})
}
