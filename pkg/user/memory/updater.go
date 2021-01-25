package memory

import (
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
	"github.com/dwethmar/atami/pkg/user/memory/util"
)

// updateRepository reads messages from memory
type updateRepository struct {
	store *memstore.Store
}

// FindAll get multiple messages
func (f updateRepository) Update(ID int, action user.UpdateAction) (*user.User, error) {
	var usr user.User

	users := f.store.GetUsers()
	if r, ok := users.Get(ID); ok {
		usr = util.FromMemory(r)

		usr.Biography = action.Biography
		usr.UpdatedAt = action.UpdatedAt

		users.Put(usr.ID, util.ToMemory(usr))
	} else {
		return nil, user.ErrCouldNotFind
	}

	return &usr, nil
}

// NewUpdater return a new in memory listin repository
func NewUpdater(
	store *memstore.Store,
) *user.Updater {
	return user.NewUpdater(
		&updateRepository{store},
	)
}
