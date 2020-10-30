package message

import (
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
)

func TestCreate(t *testing.T) {
	testCreator(t, NewMemCreator(memstore.NewMemstore()), NewMessage{
		Content: "wow",
	})
}
