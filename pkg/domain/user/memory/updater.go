package memory

import (
	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/dwethmar/atami/pkg/domain/user/memory/util"
	"github.com/dwethmar/atami/pkg/memstore"
)

// updateRepository reads messages from memory
type updateRepository struct {
	store *memstore.Memstore
}

// FindAll get multiple messages
func (f updateRepository) Update(ID int, action user.UpdateUser) (*user.User, error) {
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
	store *memstore.Memstore,
) *user.Updater {
	return user.NewUpdater(
		&updateRepository{store},
	)
}
