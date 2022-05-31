package domain

import (
	bogus "github.com/rgfaber/go-vesca/greenhouse/model"
	testing2 "github.com/rgfaber/go-vesca/sdk/core/mocks"
	"github.com/rgfaber/go-vesca/sdk/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCmd(t *testing.T) {
	id := bogus.BogusGreenhouseID
	c := NewCmd(id, testing2.TEST_TRACE_ID,
		bogus.BogusFan)
	assert.NotNil(t, c)
}

func TestCmd_ImplementsICmd(t *testing.T) {
	// Given
	cmd := NewCmd(nil, "", nil)
	// When
	ok := domain.ImplementsICmd(cmd)
	// Then
	assert.True(t, ok)
}
