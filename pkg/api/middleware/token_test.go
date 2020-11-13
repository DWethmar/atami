package middleware

import (
	"net/http/httptest"
	"testing"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticated(t *testing.T) {
	UID := auth.UID("abc")

	req := httptest.NewRequest("POST", "/", nil)
	ctx := req.Context()
	ctx = WithUserUID(ctx, UID)

	UIDFromContext, err := UserUIDFromContext(ctx)
	assert.NoError(t, err)
	assert.Equal(t, UID, UIDFromContext)
}
