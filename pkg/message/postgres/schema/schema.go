package schema

import "fmt"

var (
	// Table is the database tablename
	Table = "message"
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
	WithTbl(ColID),
	WithTbl(ColUID),
	WithTbl(ColText),
	WithTbl(ColCreatedByUserID),
	WithTbl(ColCreatedAt),
}

// WithTbl adds table to col
func WithTbl(col string) string {
	return fmt.Sprintf("%s.%s", Table, col)
}
