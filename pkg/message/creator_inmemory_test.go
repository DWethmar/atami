package message

import (
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
)

func TestCreate(t *testing.T) {
	testCreator(t, NewInMemoryCreator(memstore.NewMemstore()), NewMessage{
		Content: "wow",
	})
}
