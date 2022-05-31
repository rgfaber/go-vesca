package domain

import (
	"github.com/rgfaber/go-vesca/sdk/domain"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/mediator"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAggregate(t *testing.T) {
	// Given
	s := store.SingletonStore()
	b := mediator.SingletonDECBus()
	// When
	var a domain.IAggregate = NewAggregate(s, b)
	// Then
	assert.NotNil(t, a)
}
