package util

import (
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/node"
	"github.com/dwethmar/atami/pkg/user"
)

// ToMsgUser from the memstore to node user
func ToMsgUser(user user.User) *node.User {
	return &node.User{
		ID:       user.ID,
		UID:      user.UID,
		Username: user.Username,
	}
}

// ToMemory maps a node to memory
func ToMemory(m node.Node) memstore.Node {
	return memstore.Node{
		ID:              m.ID,
		UID:             m.UID,
		Text:            m.Text,
		CreatedByUserID: m.CreatedByUserID,
		CreatedAt:       m.CreatedAt,
	}
}

// FromMemory maps a node from memory
func FromMemory(m memstore.Node) node.Node {
	return node.Node{
		ID:              m.ID,
		UID:             m.UID,
		Text:            m.Text,
		CreatedByUserID: m.CreatedByUserID,
		CreatedAt:       m.CreatedAt,
	}
}
