package schema

import "fmt"

var (
	// Table is the database tablename
	Table = "public.message"
)

var (
	// ColID table Column
	ColID = "id"
	// ColUID table Column
	ColUID = "uid"
	// ColText table Column
	ColText = "text"
	// ColCreatedByUserID table Column
	ColCreatedByUserID = "created_by_user_id"
	// ColCreatedAt table Column
	ColCreatedAt = "created_at"
)

// SelectCols are the default selected columns
var SelectCols = []string{
	fmt.Sprintf("message.%s", ColID),
	fmt.Sprintf("message.%s", ColUID),
	fmt.Sprintf("message.%s", ColText),
	fmt.Sprintf("message.%s", ColCreatedByUserID),
	fmt.Sprintf("message.%s", ColCreatedAt),
}
