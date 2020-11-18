package schema

import "fmt"

var (
	// Table is the database tablename
	Table = "public.user"
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
	fmt.Sprintf("message.%s", ColID),
	fmt.Sprintf("message.%s", ColUID),
	fmt.Sprintf("message.%s", ColUsername),
	fmt.Sprintf("message.%s", ColEmail),
	fmt.Sprintf("message.%s", ColPassword),
	fmt.Sprintf("message.%s", ColCreatedAt),
	fmt.Sprintf("message.%s", ColUpdatedAt),
}
