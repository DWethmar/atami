package entity

import (
	"github.com/segmentio/ksuid"
)

// ID type used for the ID
type ID = int

// UID type used for the UID
type UID = string

// NewUID creates a new ID
func NewUID() UID {
	k := ksuid.New()
	return UID(k.String())
}

// StringToUID string to UID
func StringToUID(s string) (UID, error) {
	id, err := ksuid.Parse(s)
	return UID(id.String()), err
}
