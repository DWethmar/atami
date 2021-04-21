package test

import (
	"testing"

	"github.com/dwethmar/atami/pkg/user"
	"github.com/stretchr/testify/assert"
)

// TestDelete tests the Delete function.
func TestDelete(t *testing.T, deleter *user.Deleter, ID int) {
	assert.NoError(t, deleter.Delete(ID))
	assert.Equal(t, user.ErrCouldNotDelete, deleter.Delete(ID))
}
