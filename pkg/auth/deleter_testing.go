package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDelete tests the Delete function.
func TestDelete(t *testing.T, finder *Deleter, ID ID) {
	assert.Nil(t, finder.Delete(ID))
	assert.Equal(t, ErrCouldNotDelete, finder.Delete(ID))
}
