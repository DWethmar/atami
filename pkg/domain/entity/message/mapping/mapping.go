package mapping

import (
	"github.com/dwethmar/atami/pkg/domain/entity/user"
	"github.com/dwethmar/atami/pkg/memstore"
)

// UserToMemoryMap maps a message from memory
func UserToMemoryMap(m user.User) memstore.User {
	return memstore.User{
		ID:       m.ID,
		UID:      m.UID,
		Username: m.Username,
	}
}
