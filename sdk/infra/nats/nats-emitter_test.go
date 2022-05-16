package nats

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	nats "github.com/rgfaber/go-vesca/sdk/mocks/infra/nats"
	mocks_interfaces "github.com/rgfaber/go-vesca/sdk/mocks/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_New_NatsEmitter(t *testing.T) {
	//// Given
	nats := nats.NewINatsBus(t)
	eh := mocks_interfaces.NewIEventHandler(t)
	// When
	fe := NewNatsEmitter(nats, "me", eh, nil, nil)
	// Then
	assert.NotNil(t, fe)
}

func Test_NatsEmitter_Implements_IFactEmitter(t *testing.T) {
	// Given
	ne := NewNatsEmitter(nil, "", nil, nil, nil)
	// When
	ok := interfaces.ImplementsIFactEmitter(ne)
	// Then
	assert.True(t, ok)
}
