package schema

import "fmt"

var (
	// Table is the database tablename
	Table = "public.users"
)

var (
	// ColID table Column
	ColID = "id"
	// ColUID table Column
	ColUID = "uid"
	// ColUsername table Column
	ColUsername = "username"
	// ColEmail table Column
	ColEmail = "email"
	// ColPassword table Column
	ColPassword = "password"
	// ColCreatedAt table Column
	ColCreatedAt = "created_at"
	// ColUpdatedAt table Column
	ColUpdatedAt = "updated_at"
)

// SelectCols are the default selected columns
var SelectCols = []string{
	fmt.Sprintf("users.%s", ColID),
	fmt.Sprintf("users.%s", ColUID),
	fmt.Sprintf("users.%s", ColUsername),
	fmt.Sprintf("users.%s", ColEmail),
	fmt.Sprintf("users.%s", ColCreatedAt),
	fmt.Sprintf("users.%s", ColUpdatedAt),
}
