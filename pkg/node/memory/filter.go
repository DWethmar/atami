package memory

import (
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/node"
	"github.com/dwethmar/atami/pkg/node/memory/util"
)

func filterList(list []memstore.Node, filterFn func(node.Node) bool) (*node.Node, error) {
	for _, item := range list {
		node := util.FromMemory(item)
		if filterFn(node) {
			return &node, nil
		}
	}
	return nil, node.ErrCouldNotFind
}
