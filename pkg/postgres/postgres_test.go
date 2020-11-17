package postgres

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Select returns a select sql query
func TestSelect(t *testing.T) {
	expect :=
		`SELECT 
	public.user.id, 
	username
FROM public.message
LEFT JOIN public.user ON public.user.id = public.message.created_by_user_id
WHERE public.user.id >= 0
OR username LIKE '%e%'
AND public.user.created_on > '2014-02-01'
GROUP BY public.user.id, public.user.created_on
HAVING username LIKE '%e%'
ORDER BY id ASC, username DESC
LIMIT 10
OFFSET 1`

	received := Select(
		SelectQuery{
			Cols: []string{"public.user.id", "username"},
			From: "public.message",
			Joins: NewJoin().
				Left("public.user ON public.user.id = public.message.created_by_user_id"),
			Where: NewWhere().
				And("public.user.id >= 0").
				Or("username LIKE '%e%'").
				And("public.user.created_on > '2014-02-01'"),
			GroupBy: []string{"public.user.id", "public.user.created_on"},
			Having: NewWhere().
				And("username LIKE '%e%'"),
			OrderBy: []string{"id ASC", "username DESC"},
			Limit:   10,
			Offset:  1,
		},
	)

	assert.Equal(t, expect, received, fmt.Sprint(received))
}
