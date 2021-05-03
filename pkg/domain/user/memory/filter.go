package memory

import (
	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/dwethmar/atami/pkg/domain/user/memory/util"
	"github.com/dwethmar/atami/pkg/memstore"
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
