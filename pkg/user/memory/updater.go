package memory

import (
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
)

// updateRepository reads messages from memory
type updateRepository struct {
	store *memstore.Store
}

// FindAll get multiple messages
func (f updateRepository) Update(updateUser user.UpdateAction) (*user.User, error) {
	return nil, nil
}

// NewUpdater return a new in memory listin repository
func NewUpdater(
	store *memstore.Store,
) *user.Updater {
	return user.NewUpdater(
		&updateRepository{store},
	)
}
