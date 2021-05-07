package message

import (
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/domain"
	"github.com/dwethmar/atami/pkg/domain/entity"
	"github.com/stretchr/testify/assert"
)

// FindByID tests the ReadOne function.
func testFindByID(t *testing.T, reader Reader, ID entity.ID, message Message) {
	m, err := reader.Get(ID)
	assert.NoError(t, err)

	assert.NotEmpty(t, message.ID)
	assert.NotEmpty(t, message.UID)
	assert.NotEmpty(t, message.Text)
	assert.NotEmpty(t, message.CreatedByUserID)
	assert.False(t, message.CreatedAt.IsZero())

	if assert.NotNil(t, m) {
		assert.NotEmpty(t, m.ID)
		assert.NotEmpty(t, m.UID)
		assert.NotEmpty(t, m.Text)
		assert.NotEmpty(t, m.CreatedByUserID)
		assert.False(t, m.CreatedAt.IsZero())

		assert.Equal(t, message.ID, m.ID)
		assert.Equal(t, message.Text, m.Text)
		assert.Equal(t, message.CreatedByUserID, m.CreatedByUserID)

		if assert.NotNil(t, m.User) {
			assert.Equal(t, message.CreatedByUserID, m.User.ID)
			assert.Equal(t, m.CreatedByUserID, m.User.ID)
		}
	}
}

// NotFoundByID tests the ReadOne function for a not found error.
func testNotFoundByID(t *testing.T, reader Reader) {
	_, err := reader.Get(1)
	assert.Equal(t, domain.ErrNotFound, err)
}

// FindByUID tests the findByUID function.
func testFindByUID(t *testing.T, reader Reader, UID entity.UID, message Message) {
	m, err := reader.GetByUID(UID)
	assert.NoError(t, err)

	assert.NotEmpty(t, message.ID)
	assert.NotEmpty(t, message.UID)
	assert.NotEmpty(t, message.Text)
	assert.NotEmpty(t, message.CreatedByUserID)
	assert.False(t, message.CreatedAt.IsZero())

	if assert.NotNil(t, m) {
		assert.NotEmpty(t, m.ID)
		assert.NotEmpty(t, m.UID)
		assert.NotEmpty(t, m.Text)
		assert.NotEmpty(t, m.CreatedByUserID)
		assert.False(t, m.CreatedAt.IsZero())

		assert.Equal(t, message.ID, m.ID)
		assert.Equal(t, message.Text, m.Text)
		assert.Equal(t, message.CreatedByUserID, m.CreatedByUserID)

		if assert.NotNil(t, m.User) {
			assert.Equal(t, message.CreatedByUserID, m.User.ID)
			assert.Equal(t, m.CreatedByUserID, m.User.ID)
		}
	}
}

// NotFoundByUID tests the ReadOne function for a not found error.
func testNotFoundByUID(t *testing.T, reader Reader) {
	_, err := reader.GetByUID("d")
	assert.Equal(t, domain.ErrNotFound, err)
}

// Find tests the Find function.
func testFind(t *testing.T, reader Reader, length uint, messages []Message) {
	list, err := reader.List(0, length)

	assert.NoError(t, err)
	if assert.Equal(t, length, len(list)) {
		for i, message := range list {
			assert.NotEmpty(t, messages[i].ID)
			assert.NotEmpty(t, messages[i].UID)
			assert.NotEmpty(t, messages[i].Text)
			assert.NotZero(t, messages[i].CreatedByUserID)
			assert.False(t, messages[i].CreatedAt.IsZero())

			assert.Equal(t, messages[i].ID, message.ID)
			assert.Equal(t, messages[i].UID, message.UID)
			assert.Equal(t, messages[i].Text, message.Text)

			if assert.NotNil(t, message.User) {
				assert.Equal(t, message.CreatedByUserID, message.User.ID)
			}
		}
	}
}

// Store test the writer store
func testStore(t *testing.T, writer Writer, reader Reader, create Create) {
	ID, err := writer.Create(create)
	message, err := reader.Get(ID)

	assert.Nil(t, err)
	assert.Equal(t, message.ID, 1)
	assert.Equal(t, message.Text, create.Text)
	assert.Equal(t, message.CreatedByUserID, create.CreatedByUserID)
	assert.True(t, time.Now().Add(time.Microsecond).After(message.CreatedAt))
}

// InvalidCreate tests an invalid create
func testInvalidCreate(t *testing.T, writer Writer, create Create) {
	_, err := writer.Create(create)
	assert.Error(t, err)
}

// Delete tests the Delete function.
func testDelete(t *testing.T, writer Writer, ID entity.ID) {
	assert.Nil(t, writer.Delete(ID))
	assert.Equal(t, domain.ErrCannotBeDeleted, writer.Delete(ID))
}
