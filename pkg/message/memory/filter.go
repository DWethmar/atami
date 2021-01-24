package memory

import (
	"github.com/dwethmar/atami/pkg/message"
	"github.com/dwethmar/atami/pkg/user"
)

func filterList(list []interface{}, filterFn func(message.Message) bool) (*message.Message, error) {
	for _, item := range list {
		if record, ok := item.(message.Message); ok {
			if filterFn(record) {
				return &record, nil
			}
		} else {
			return nil, errCouldNotParse
		}
	}
	return nil, user.ErrCouldNotFind
}
