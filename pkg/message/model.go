package message

import (
	"github.com/dwethmar/atami/pkg/models"
	"github.com/dwethmar/atami/pkg/types"
)

// Message model
type Message struct {
	models.Message
}

func (m *Message) GetID() types.ID {
	return m.ID
}
