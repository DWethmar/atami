package message

import (
	"testing"

	"github.com/dwethmar/atami/pkg/model"
	"github.com/stretchr/testify/assert"
)

// TestDelete tests the Delete function.
func TestDelete(t *testing.T, repo DeleterRepository, ID model.MessageID) {
	assert.Nil(t, repo.Delete(ID))
	assert.Equal(t, ErrCouldNotDelete, repo.Delete(ID))
}
