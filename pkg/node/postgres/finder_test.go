package postgres

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/node"
	"github.com/stretchr/testify/assert"
)

func generateTestNodes(size int) []node.CreateAction {
	nodes := make([]node.CreateAction, size)
	for i := 0; i < size; i++ {
		nodes[i] = node.CreateAction{
			UID:             fmt.Sprintf("%v", i),
			Text:            fmt.Sprintf("Lorum ipsum %d", i+1),
			CreatedByUserID: 1,
			CreatedAt:       time.Now().AddDate(0, -1, 0).Add(time.Duration(i) * time.Second),
		}
	}
	return nodes
}

func setup(db *sql.DB, size int) (*node.Finder, []node.Node) {
	nodes := make([]node.Node, size)

	createRepo := &creatorRepository{db}
	findRepo := &findRepository{db}

	for i, newMSG := range generateTestNodes(size) {
		if msg, err := createRepo.Create(newMSG); err == nil {
			m, _ := findRepo.FindByID(msg.ID)
			nodes[i] = *m
		} else {
			fmt.Printf("error: %s", err)
			panic(1)
		}
	}
	return NewFinder(db), nodes
}

func TestFindByID(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		finder, nodes := setup(db, 100)
		node.TestFindByID(t, finder, 10, nodes[9])
		return nil
	}))
}

func TestNotFound(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		finder, _ := setup(db, 100)
		node.TestNotFound(t, finder)
		return nil
	}))
}

func TestFind(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		finder, m := setup(db, 300)
		items := m[250:300]

		// Reverse items because of the order by on created_at DESC
		for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
			items[i], items[j] = items[j], items[i]
		}

		node.TestFind(t, finder, 50, items)
		return nil
	}))
}
