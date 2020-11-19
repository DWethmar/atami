package querybuilder

import (
	"fmt"
	"strings"
)

// UpdateQuery defines the fields to build a update sql query
type UpdateQuery struct {
	Table     string
	Set       map[string]interface{}
	From      *SelectQuery
	FromAs    string
	Where     *Where
	Returning []string
}

// Update returns a update sql query
func Update(uq UpdateQuery) string {
	queryParts := []string{}

	if uq.Table != "" {
		fromPart := fmt.Sprintf(`UPDATE %s`, uq.Table)
		queryParts = append(queryParts, fromPart)
	}

	if len(uq.Set) > 0 {
		setParts := []string{}

		for column, v := range uq.Set {
			switch v.(type) {
			case int:
				setParts = append(setParts, fmt.Sprintf("%s = %d", column, v))
			case float32:
				setParts = append(setParts, fmt.Sprintf("%s = %d", column, v))
			case string:
				setParts = append(setParts, fmt.Sprintf("%s = %s", column, v))
			case interface{}:
				setParts = append(setParts, fmt.Sprintf("%s = %v", column, v))
			}
		}

		queryParts = append(
			queryParts,
			fmt.Sprintf(`SET %s`, "\n\t"+strings.Join(setParts, ",\n\t")+"\n"),
		)
	}

	if uq.From != nil && uq.FromAs != "" {
		queryParts = append(
			queryParts,
			fmt.Sprintf("FROM (%s) AS %s", Select(*uq.From), uq.FromAs),
		)
	}

	if uq.Where != nil {
		queryParts = append(queryParts, fmt.Sprintf(`WHERE %s`, uq.Where.String()))
	}

	if len(uq.Returning) > 0 {
		queryParts = append(
			queryParts,
			fmt.Sprintf(`RETURNING %s`, strings.Join(uq.Returning, ",")),
		)
	}

	return strings.Join(queryParts, "\n")
}
