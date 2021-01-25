package querybuilder

import (
	"fmt"
	"strings"
)

// UpdateCol specifies what updates
type UpdateCol struct {
	Name  string
	Value interface{}
}

// UpdateQuery defines the fields to build a update sql query
type UpdateQuery struct {
	Table     string
	Set       []UpdateCol
	From      *SelectQuery
	FromAs    string // alias for From
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

		for _, v := range uq.Set {
			setParts = append(setParts, fmt.Sprintf("%s = %v", v.Name, v.Value))
		}

		queryParts = append(
			queryParts,
			fmt.Sprintf(`SET%s`, "\n\t"+strings.Join(setParts, ",\n\t")),
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
			fmt.Sprintf(`RETURNING %s`, strings.Join(uq.Returning, ", ")),
		)
	}

	return strings.Join(queryParts, "\n")
}
