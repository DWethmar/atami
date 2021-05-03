package validate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var biographyValidator = NewBiographyValidator()

func TestBiography(t *testing.T) {
	valid := []string{
		"lorum ipsum",
		"",
		"a",
	}

	for _, v := range valid {
		assert.NoError(t, biographyValidator.Validate(v), fmt.Sprintf("Biography: %s", v))
	}
}

func TestInvalidBiography(t *testing.T) {
	valid := []string{
		"lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum",
	}

	for _, v := range valid {
		assert.Error(t, biographyValidator.Validate(v), fmt.Sprintf("Biography: %s", v))
	}
}
