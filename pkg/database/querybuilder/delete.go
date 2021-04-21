package querybuilder

import (
	"fmt"
	"strings"
)

// DeleteQuery defines the fields to build a delete sql query
type DeleteQuery struct {
	From      From
	Where     *Where
	Returning []string
}

// Delete returns a delete sql query
func Delete(dq DeleteQuery) string {
	queryParts := []string{}

	if dq.From != "" {
		fromPart := fmt.Sprintf(`DELETE %s`, dq.From.String())
		queryParts = append(queryParts, fromPart)
	}

	if dq.Where != nil {
		queryParts = append(queryParts, fmt.Sprintf(`WHERE %s`, dq.Where.String()))
	}

	if len(dq.Returning) > 0 {
		queryParts = append(
			queryParts,
			fmt.Sprintf(`RETURNING %s`, strings.Join(dq.Returning, ",")),
		)
	}

	return strings.Join(queryParts, "\n")
}
