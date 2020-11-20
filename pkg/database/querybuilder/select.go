package querybuilder

import (
	"fmt"
	"strings"
)

// SelectQuery defines the fields to build a sql query
type SelectQuery struct {
	Cols    []string
	From    string
	Joins   *Join
	Where   *Where
	GroupBy []string
	Having  *Where
	OrderBy []string
	Limit   string
	Offset  string
}

// Select returns a select sql query
func Select(sq SelectQuery) string {
	queryParts := []string{}

	if len(sq.Cols) > 0 {
		selectPart := fmt.Sprintf(`SELECT%s`, "\n\t"+strings.Join(sq.Cols, ",\n\t"))
		queryParts = append(queryParts, selectPart)
	} else {
		queryParts = append(queryParts, "SELECT *")
	}

	if sq.From != "" {
		fromPart := fmt.Sprintf(`FROM %s`, sq.From)
		queryParts = append(queryParts, fromPart)
	}

	if sq.Joins != nil {
		queryParts = append(queryParts, sq.Joins.String())
	}

	if sq.Where != nil {
		queryParts = append(queryParts, fmt.Sprintf(`WHERE %s`, sq.Where.String()))
	}

	if len(sq.GroupBy) > 0 {
		groupByPart := fmt.Sprintf(`GROUP BY %s`, strings.Join(sq.GroupBy, ", "))
		queryParts = append(queryParts, groupByPart)
	}

	if sq.Having != nil {
		queryParts = append(queryParts, fmt.Sprintf(`HAVING %s`, sq.Having.String()))
	}

	if len(sq.OrderBy) > 0 {
		orderByPart := fmt.Sprintf(`ORDER BY %s`, strings.Join(sq.OrderBy, ", "))
		queryParts = append(queryParts, orderByPart)
	}

	if sq.Limit != "" {
		queryParts = append(queryParts, fmt.Sprintf(`LIMIT %s`, sq.Limit))
	}

	if sq.Offset != "" {
		queryParts = append(queryParts, fmt.Sprintf(`OFFSET %s`, sq.Offset))
	}

	return strings.Join(queryParts, "\n")
}
