package querybuilder

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	expect := `INSERT INTO public.user
(
	uid,
	email,
	username,
	password,
	created_at,
	updated_at
)
VALUES (
	'abcd123',
	'test@test.nl',
	'username',
	'p@ssw0rd',
	'2020-11-13 15:33:00.972651',
	'2020-11-13 15:33:00.972651'
)`

	received := Insert(
		InsertQuery{
			Into: "public.user",
			Cols: []string{
				"uid",
				"email",
				"username",
				"password",
				"created_at",
				"updated_at",
			},
			Values: []interface{}{
				"abcd123",
				"test@test.nl",
				"username",
				"p@ssw0rd",
				"2020-11-13 15:33:00.972651",
				"2020-11-13 15:33:00.972651",
			},
		},
	)

	assert.Equal(t, expect, received, fmt.Sprint(received))
}

func TestInsertWithSubquery(t *testing.T) {
	expect := `INSERT INTO public.user
(
	uid,
	email,
	username,
	password,
	created_at,
	updated_at
)
SELECT
	'abcdefg',
	CONCAT('c_', email),
	CONCAT('c_', username),
	'passwordlmasxjlkasjd',
	created_at,
	updated_at
FROM public.user
WHERE id = 1`

	received := Insert(
		InsertQuery{
			Into: "public.user",
			Cols: []string{
				"uid",
				"email",
				"username",
				"password",
				"created_at",
				"updated_at",
			},
			Select: &SelectQuery{
				Cols: []string{
					"'abcdefg'",
					"CONCAT('c_', email)",
					"CONCAT('c_', username)",
					"'passwordlmasxjlkasjd'",
					"created_at",
					"updated_at",
				},
				From:  "public.user",
				Where: NewWhere().And("id = 1"),
			},
		},
	)

	assert.Equal(t, expect, received, fmt.Sprint(received))
}
