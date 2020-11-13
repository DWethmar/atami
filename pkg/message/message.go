package message

import (
	"time"

	"github.com/dwethmar/atami/pkg/model"
)

// The Message model
type Message struct {
	ID        model.MessageID
	UID       model.MessageUID
	Text      string
	CreatedAt time.Time
}

func toMessage(message *Message) *model.Message {
	return &model.Message{
		ID:        message.ID,
		UID:       message.UID,
		Text:      message.Text,
		CreatedAt: message.CreatedAt,
	}
}
