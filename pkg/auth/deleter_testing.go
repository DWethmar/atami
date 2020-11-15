package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDelete tests the Delete function.
func TestDelete(t *testing.T, deleter *Deleter, ID int) {
	assert.Nil(t, deleter.Delete(ID))
	assert.Equal(t, ErrCouldNotDelete, deleter.Delete(ID))
}
