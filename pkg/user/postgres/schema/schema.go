package schema

var (
	// Table is the database tablename
	Table = "public.app_user"
)

const (
	// ColID table Column
	ColID = "app_user.id"
	// ColUID table Column
	ColUID = "app_user.uid"
	// ColUsername table Column
	ColUsername = "app_user.username"
	// ColEmail table Column
	ColEmail = "app_user.email"
	// ColPassword table Column
	ColPassword = "app_user.password"
	// ColCreatedAt table Column
	ColCreatedAt = "app_user.created_at"
	// ColUpdatedAt table Column
	ColUpdatedAt = "app_user.updated_at"
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
