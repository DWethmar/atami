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
	CreatedBy model.UserID
	CreatedAt time.Time
}

// ToMessage maps a message to a model message
func ToMessage(message *Message) *model.Message {
	return &model.Message{
		ID:        message.ID,
		UID:       message.UID,
		Text:      message.Text,
		CreatedBy: message.CreatedBy,
		CreatedAt: message.CreatedAt,
	}
}
