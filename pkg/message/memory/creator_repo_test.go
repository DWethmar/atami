package memory

import (
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	newMessage := message.NewMessage{
		Content: "wow",
	}

	repo := NewCreatorRepository(memstore.New())

	m, err := repo.Create(newMessage)
	assert.Nil(t, err)
	assert.Equal(t, m.ID, message.ID(1))
	assert.Equal(t, m.Content, newMessage.Content)
	time.Sleep(1)
	assert.True(t, time.Now().After(m.CreatedAt))
}
