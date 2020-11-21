package querybuilder

import (
	"fmt"
	"strings"
)

// InsertQuery defines the fields to build a insert sql query
type InsertQuery struct {
	Into      string
	Cols      []string
	Values    []interface{}
	Select    *SelectQuery // If populated than instead of values, a select will be print
	Returning []string
}

// Insert returns a insert sql query
func Insert(iq InsertQuery) string {
	queryParts := []string{}

	if iq.Into != "" {
		fromPart := fmt.Sprintf(`INSERT INTO %s`, iq.Into)
		queryParts = append(queryParts, fromPart)
	}

	if len(iq.Cols) > 0 {
		selectPart := fmt.Sprintf(`(%s)`, "\n\t"+strings.Join(iq.Cols, ",\n\t")+"\n")
		queryParts = append(queryParts, selectPart)
	}

	if len(iq.Values) > 0 && iq.Select == nil {
		valuesParts := []string{}

		for _, v := range iq.Values {
			valuesParts = append(valuesParts, fmt.Sprintf("%v", v))
		}

		queryParts = append(
			queryParts,
			fmt.Sprintf(`VALUES (%s)`, "\n\t"+strings.Join(valuesParts, ",\n\t")+"\n"),
		)
	}

	if iq.Select != nil {
		queryParts = append(queryParts, Select(*iq.Select))
	}

	if len(iq.Returning) > 0 {
		queryParts = append(
			queryParts,
			fmt.Sprintf(`RETURNING %s`, strings.Join(iq.Returning, ", ")),
		)
	}

	return strings.Join(queryParts, "\n")
}
