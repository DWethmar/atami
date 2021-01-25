package node

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFindByID tests the ReadOne function.
func TestFindByID(t *testing.T, finder *Finder, ID int, node Node) {
	m, err := finder.FindByID(ID)
	assert.NoError(t, err)

	assert.NotEmpty(t, node.ID)
	assert.NotEmpty(t, node.UID)
	assert.NotEmpty(t, node.Text)
	assert.NotEmpty(t, node.CreatedByUserID)
	assert.False(t, node.CreatedAt.IsZero())

	if assert.NotNil(t, m) {
		assert.NotEmpty(t, m.ID)
		assert.NotEmpty(t, m.UID)
		assert.NotEmpty(t, m.Text)
		assert.NotEmpty(t, m.CreatedByUserID)
		assert.False(t, m.CreatedAt.IsZero())

		assert.Equal(t, node.ID, m.ID)
		assert.Equal(t, node.Text, m.Text)
		assert.Equal(t, node.CreatedByUserID, m.CreatedByUserID)

		if assert.NotNil(t, m.User) {
			assert.Equal(t, node.CreatedByUserID, m.User.ID)
			assert.Equal(t, m.CreatedByUserID, m.User.ID)
		}
	}
}

// TestFindByUID tests the findByUID function.
func TestFindByUID(t *testing.T, finder *Finder, UID string, node Node) {
	m, err := finder.FindByUID(UID)
	assert.NoError(t, err)

	assert.NotEmpty(t, node.ID)
	assert.NotEmpty(t, node.UID)
	assert.NotEmpty(t, node.Text)
	assert.NotEmpty(t, node.CreatedByUserID)
	assert.False(t, node.CreatedAt.IsZero())

	if assert.NotNil(t, m) {
		assert.NotEmpty(t, m.ID)
		assert.NotEmpty(t, m.UID)
		assert.NotEmpty(t, m.Text)
		assert.NotEmpty(t, m.CreatedByUserID)
		assert.False(t, m.CreatedAt.IsZero())

		assert.Equal(t, node.ID, m.ID)
		assert.Equal(t, node.Text, m.Text)
		assert.Equal(t, node.CreatedByUserID, m.CreatedByUserID)

		if assert.NotNil(t, m.User) {
			assert.Equal(t, node.CreatedByUserID, m.User.ID)
			assert.Equal(t, m.CreatedByUserID, m.User.ID)
		}
	}
}

// TestNotFound tests the ReadOne function for a not found error.
func TestNotFound(t *testing.T, finder *Finder) {
	_, err := finder.FindByID(0)
	assert.Equal(t, ErrCouldNotFind, err)
}

// TestFind tests the Find function.
func TestFind(t *testing.T, finder *Finder, length int, nodes []Node) {
	list, err := finder.Find(0, length)

	assert.NoError(t, err)
	if assert.Equal(t, length, len(list)) {
		for i, node := range list {
			assert.NotEmpty(t, nodes[i].ID)
			assert.NotEmpty(t, nodes[i].UID)
			assert.NotEmpty(t, nodes[i].Text)
			assert.NotZero(t, nodes[i].CreatedByUserID)
			assert.False(t, nodes[i].CreatedAt.IsZero())

			assert.Equal(t, nodes[i].ID, node.ID)
			assert.Equal(t, nodes[i].UID, node.UID)
			assert.Equal(t, nodes[i].Text, node.Text)

			if assert.NotNil(t, node.User) {
				// fmt.Println(node)
				// fmt.Println("UID: -> " + node.User.UID + " <- ")
				// fmt.Println(node.User.ID)
				// fmt.Println(node.User.Username)
				assert.Equal(t, node.CreatedByUserID, node.User.ID)
			}
		}
	}
}
