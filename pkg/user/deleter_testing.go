package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDelete tests the Delete function.
func TestDelete(t *testing.T, repo DeleterRepository, ID ID) {
	assert.Nil(t, repo.Delete(ID))
	assert.Equal(t, ErrCouldNotDelete, repo.Delete(ID))
}
