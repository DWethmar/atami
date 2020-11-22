package querybuilder

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	expect := `DELETE FROM public.users
WHERE id = 1`

	received := Delete(
		DeleteQuery{
			From:  "public.users",
			Where: NewWhere().And("id = 1"),
		},
	)

	assert.Equal(t, expect, received, fmt.Sprint(received))
}
