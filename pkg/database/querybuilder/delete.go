package querybuilder

import (
	"fmt"
	"strings"
)

// DeleteQuery defines the fields to build a delete sql query
type DeleteQuery struct {
	From  string
	Where *Where
}

// Delete returns a delete sql query
func Delete(dq DeleteQuery) string {
	queryParts := []string{}

	if dq.From != "" {
		fromPart := fmt.Sprintf(`DELETE FROM %s`, dq.From)
		queryParts = append(queryParts, fromPart)
	}

	if dq.Where != nil {
		queryParts = append(queryParts, fmt.Sprintf(`WHERE %s`, dq.Where.String()))
	}

	return strings.Join(queryParts, "\n")
}
