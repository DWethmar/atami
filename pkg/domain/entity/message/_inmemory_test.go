package message

import (
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/stretchr/testify/assert"
)

func TestMemoryFindByID(t *testing.T) {
	memStore := memstore.NewStore()
	memStore.GetUsers().Put(1, *memstore.NewFixtureUser(1))

	repo := NewinMemoryRepoRepository(memStore)
	ID, err := repo.Create(Create{
		UID:             "abc",
		Text:            "asd",
		CreatedByUserID: 1,
	})
	if assert.NoError(t, err) {
		testFindByID(t, repo, ID, Message{
			ID:              1,
			UID:             "abc",
			Text:            "asd",
			CreatedByUserID: 1,
		})
	}
}

func TestMemoryNotFindByID(t *testing.T) {
	repo := NewinMemoryRepoRepository(memstore.NewStore())
	testNotFoundByID(t, repo)
}

func TestMemoryFindByUID(t *testing.T) {

}

func TestMemoryFindAll(t *testing.T) {

}

func TestMemoryUpdate(t *testing.T) {
}

func TestMemoryStore(t *testing.T) {

}

func TestMemoryDelete(t *testing.T) {

}
