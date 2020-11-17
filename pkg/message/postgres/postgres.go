package postgres

import "errors"

//go:generate go run ./generate/gen.go

var (
	// Table is the database tablename
	Table            = "public.message"
	errCouldNotParse = errors.New("could not parse user")
)

var (
	// ColID table Column
	ColID = "message.id"
	// ColUID table Column
	ColUID = "message.uid"
	// Coltext table Column
	Coltext = "message.text"
	// ColCreatedByUserID table Column
	ColCreatedByUserID = "message.created_by_user_id"
	// ColCreatedAt table Column
	ColCreatedAt = "message.created_at"
)

// SelectCols are the default selected columns
var SelectCols = []string{
	ColID,
	ColUID,
	Coltext,
	ColCreatedByUserID,
	ColCreatedAt,
}
