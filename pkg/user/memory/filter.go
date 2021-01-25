package memory

import (
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
	"github.com/dwethmar/atami/pkg/user/memory/util"
)

func filterList(list []memstore.User, filterFunc func(user.User) bool) (*user.User, error) {
	for _, item := range list {
		user := util.FromMemory(item)
		if filterFunc(user) {
			return &user, nil
		}
	}
	return nil, user.ErrCouldNotFind
}
