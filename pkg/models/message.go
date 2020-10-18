package models

import (
	"time"

	"github.com/dwethmar/atami/pkg/types"
)

// Message model
type Message struct {
	ID        types.ID
	Content   string
	CreatedOn time.Time
}
