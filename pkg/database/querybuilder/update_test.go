package querybuilder

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	expect := `UPDATE public.user
SET
	uid = 'abcd0987',
	email = 'updated_test@test.com',
	username = 'updated_usr',
	password = 'p@ssw0rd',
	created_at = '2020-11-13 15:33:00.972651',
	updated_at = '2020-11-13 15:33:00.972651'
)
WHERE id = 1
RETURNING id`

	received := Update(
		UpdateQuery{
			Table: "public.user",
			Set: map[string]interface{}{
				"uid":        "'abcd0987'",
				"email":      "'updated_test@test.com'",
				"username":   "'updated_usr'",
				"password":   "'p@ssw0rd'",
				"created_at": "'2020-11-13 15:33:00.972651'",
				"updated_at": "'2020-11-13 15:33:00.972651'",
			},
			Where:     NewWhere().And("id = 1"),
			Returning: []string{"id"},
		},
	)

	assert.Equal(t, expect, received, fmt.Sprint(received))
}

// func TestUpdateWithSubquery(t *testing.T) {
// 	expect := `INSERT INTO public.user
// (
// 	uid,
// 	email,
// 	username,
// 	password,
// 	created_at,
// 	updated_at
// )
// SELECT
// 	'abcdefg',
// 	CONCAT('c_', email),
// 	CONCAT('c_', username),
// 	'passwordlmasxjlkasjd',
// 	created_at,
// 	updated_at
// FROM public.user
// WHERE id = 1
// RETURNING id`

// 	received := Insert(
// 		InsertQuery{
// 			Into: "public.user",
// 			Cols: []string{
// 				"uid",
// 				"email",
// 				"username",
// 				"password",
// 				"created_at",
// 				"updated_at",
// 			},
// 			Select: &SelectQuery{
// 				Cols: []string{
// 					"'abcdefg'",
// 					"CONCAT('c_', email)",
// 					"CONCAT('c_', username)",
// 					"'passwordlmasxjlkasjd'",
// 					"created_at",
// 					"updated_at",
// 				},
// 				From:  "public.user",
// 				Where: NewWhere().And("id = 1"),
// 			},
// 			Returning: []string{"id"},
// 		},
// 	)

// 	assert.Equal(t, expect, received, fmt.Sprint(received))
// }
