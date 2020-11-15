package memory

import "github.com/dwethmar/atami/pkg/user"

func filterList(list []interface{}, filterFn func(userRecord) bool) (*user.User, error) {
	for _, item := range list {
		if record, ok := item.(userRecord); ok {
			if filterFn(record) {
				return recordToUser(record), nil
			}
		} else {
			return nil, errCouldNotParse
		}
	}
	return nil, user.ErrCouldNotFind
}
