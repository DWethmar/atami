package schema

import "fmt"

var (
	// Table is the database tablename
	Table = "public.app_user"
)

const (
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
	WithTbl(ColID),
	WithTbl(ColUID),
	WithTbl(ColUsername),
	WithTbl(ColEmail),
	WithTbl(ColCreatedAt),
	WithTbl(ColUpdatedAt),
}

// WithTbl adds table to col
func WithTbl(col string) string {
	return fmt.Sprintf("%s.%s", Table, col)
}
