package memory

import (
	"fmt"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/node"
	"github.com/dwethmar/atami/pkg/node/memory/util"
)

func generateTestNodes(size int) []node.CreateRequest {
	nodes := make([]node.CreateRequest, size)
	for i := 0; i < size; i++ {
		nodes[i] = node.CreateRequest{
			Text:            fmt.Sprintf("Lorum ipsum %v", i+1),
			CreatedByUserID: 1,
		}
	}
	return nodes
}

func setup() (*memstore.Store, []node.Node) {
	store := memstore.NewStore()
	util.AddTestUser(store, 1)

	service := New(store)
	msgs := make([]node.Node, 100)
	for i, newMsg := range generateTestNodes(100) {
		if msg, err := service.Create(newMsg); err == nil {
			msgs[i] = *msg
		} else {
			fmt.Printf("error: %s", err)
			panic(1)
		}
	}
	return store, msgs
}

func TestByUID(t *testing.T) {
	store, nodes := setup()
	node.TestFindByUID(t, NewFinder(store), nodes[0].UID, nodes[0])
}

func TestFindByID(t *testing.T) {
	store, _ := setup()
	node.TestFindByID(t, NewFinder(store), 1, node.Node{
		ID:              1,
		UID:             "abcdef",
		Text:            "Lorum ipsum 1",
		CreatedByUserID: 1,
		CreatedAt:       time.Now(),
	})
}

func TestNotFound(t *testing.T) {
	store, _ := setup()
	node.TestNotFound(t, NewFinder(store))
}

func TestFindAll(t *testing.T) {
	store, nodes := setup()
	node.TestFind(t, NewFinder(store), 100, nodes)
}
