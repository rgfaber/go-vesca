package measure

import (
	"github.com/rgfaber/go-vesca/greenhouse/infra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAggregate(t *testing.T) {
	// Given
	s := infra.NewStore()
	b := dec.NewDECBus()
	// When
	var a dec.IAggregate = NewAggregate(s, b)
	// Then
	assert.NotNil(t, a)
}
