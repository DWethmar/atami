package schema

var (
	// Table is the database tablename
	Table = "public.message"
)

var (
	// ColID table Column
	ColID = "message.id"
	// ColUID table Column
	ColUID = "message.uid"
	// ColText table Column
	ColText = "message.text"
	// ColCreatedByUserID table Column
	ColCreatedByUserID = "message.created_by_user_id"
	// ColCreatedAt table Column
	ColCreatedAt = "message.created_at"
)

// SelectCols are the default selected columns
var SelectCols = []string{
	ColID,
	ColUID,
	ColText,
	ColCreatedByUserID,
	ColCreatedAt,
}
