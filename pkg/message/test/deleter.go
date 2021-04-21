package test

import (
	"testing"

	"github.com/dwethmar/atami/pkg/message"
	"github.com/stretchr/testify/assert"
)

// Delete tests the Delete function.
func Delete(t *testing.T, repo message.DeleterRepository, ID int) {
	assert.Nil(t, repo.Delete(ID))
	assert.Equal(t, message.ErrCouldNotDelete, repo.Delete(ID))
}
