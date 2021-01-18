package schema

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
	ColID,
	ColUID,
	ColUsername,
	ColEmail,
	ColCreatedAt,
	ColUpdatedAt,
}
