package sdk

import (
	"github.com/nauyey/guard"
	"github.com/nauyey/guard/validators"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGuard(t *testing.T) {
	// Given
	name := "User Name"
	age := 10
	gender := ""

	// When
	// validate data
	err := guard.Validate(
		&validators.StringNotBlank{Value: name},
		&validators.IntGreaterThan{Value: age, Target: 16},                                  // invalid age
		&validators.StringInclusion{Value: gender, In: []string{"female", "male", "other"}}, // invalid gender
	)
	errs, ok := err.(guard.Errors)

	// Then
	assert.Equal(t, 2, len(errs.ValidationErrors()))
	assert.True(t, ok)

}
