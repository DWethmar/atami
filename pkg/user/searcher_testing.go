package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSearchByEmail tests the search function.
func TestSearchByEmail(t *testing.T, repo SearchRepository, length int, email string) {
	r, err := repo.SearchByEmail(email)

	assert.Nil(t, err)
	assert.Equal(t, length, len(r))

	for _, user := range r {
		assert.Equal(t, email, user.Email)
	}
}
