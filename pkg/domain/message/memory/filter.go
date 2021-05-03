package memory

import (
	"github.com/dwethmar/atami/pkg/domain/message"
	"github.com/dwethmar/atami/pkg/domain/message/memory/util"
	"github.com/dwethmar/atami/pkg/memstore"
)

func filterList(list []memstore.Message, filterFn func(message.Message) bool) (*message.Message, error) {
	for _, item := range list {
		message := util.FromMemory(item)
		if filterFn(message) {
			return &message, nil
		}
	}
	return nil, message.ErrCouldNotFind
}
