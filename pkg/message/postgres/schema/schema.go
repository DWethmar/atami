package schema

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
	ColID,
	ColUID,
	ColText,
	ColCreatedByUserID,
	ColCreatedAt,
}
