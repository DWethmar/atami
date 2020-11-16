package postgres

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Select returns a select sql query
func TestSelect(t *testing.T) {
	expect :=
		`SELECT id, username
FROM public.user
WHERE id >= 0
OR username LIKE '%test%'
AND username LIKE '%account%'
GROUP BY id, username
HAVING username LIKE '%e%'
ORDER BY id ASC, username DESC
LIMIT 10
OFFSET 1`

	w := &Where{}
	w.And("id >= 0")
	w.Or("username LIKE '%test%'")
	w.And("username LIKE '%account%'")

	h := &Where{}
	h.And("username LIKE '%e%'")

	received, err := Select(
		[]string{"id", "username"},
		"public.user",
		w,
		[]string{"id", "username"},
		h,
		[]string{"id ASC", "username DESC"},
		10,
		1,
	)

	assert.NoError(t, err)
	assert.Equal(t, expect, received, fmt.Sprint(received))
}
