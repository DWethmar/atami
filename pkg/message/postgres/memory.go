package postgres

import "errors"

var (
	tableName        = "public.message"
	errCouldNotParse = errors.New("could not parse user")
)
