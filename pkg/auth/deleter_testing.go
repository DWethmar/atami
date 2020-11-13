package auth

import (
	"testing"

	"github.com/dwethmar/atami/pkg/model"
	"github.com/stretchr/testify/assert"
)

// TestDelete tests the Delete function.
func TestDelete(t *testing.T, deleter *Deleter, ID model.UserID) {
	assert.Nil(t, deleter.Delete(ID))
	assert.Equal(t, ErrCouldNotDelete, deleter.Delete(ID))
}
