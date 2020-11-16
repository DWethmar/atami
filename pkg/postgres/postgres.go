package postgres

import (
	"errors"
	"fmt"
	"strings"
)

// Select returns a select sql query
func Select(
	cols []string,
	table string,
	where *Where,
	groupBy []string,
	having *Where,
	orderBy []string,
	limit int,
	offset int,
) (string, error) {
	queryParts := []string{}

	if len(cols) == 0 {
		return "", errors.New("columns are required")
	}
	selectPart := fmt.Sprintf(`SELECT %s`, strings.Join(cols, ", "))
	queryParts = append(queryParts, selectPart)

	if table == "" {
		return "", errors.New("table is required")
	}
	fromPart := fmt.Sprintf(`FROM %s`, table)
	queryParts = append(queryParts, fromPart)

	if where != nil {
		queryParts = append(queryParts, fmt.Sprintf(`WHERE %s`, where.String()))
	}

	if len(groupBy) > 0 {
		groupByPart := fmt.Sprintf(`GROUP BY %s`, strings.Join(groupBy, ", "))
		queryParts = append(queryParts, groupByPart)
	}

	if having != nil {
		queryParts = append(queryParts, fmt.Sprintf(`HAVING %s`, having.String()))
	}

	orderByPart := fmt.Sprintf(`ORDER BY %s`, strings.Join(orderBy, ", "))
	queryParts = append(queryParts, orderByPart)

	if limit > 0 {
		queryParts = append(queryParts, fmt.Sprintf(`LIMIT %d`, limit))
	}

	if offset > 0 {
		queryParts = append(queryParts, fmt.Sprintf(`OFFSET %d`, offset))
	}

	return strings.Join(queryParts, "\n"), nil
}
