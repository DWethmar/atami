package schema

import "fmt"

var (
	// Table is the database tablename
	Table = "public.app_user"
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
	fmt.Sprintf("app_user.%s", ColID),
	fmt.Sprintf("app_user.%s", ColUID),
	fmt.Sprintf("app_user.%s", ColUsername),
	fmt.Sprintf("app_user.%s", ColEmail),
	fmt.Sprintf("app_user.%s", ColCreatedAt),
	fmt.Sprintf("app_user.%s", ColUpdatedAt),
}
