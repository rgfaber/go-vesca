package nats

import (
	"github.com/rgfaber/go-vesca/greenhouse/features/initialize/contract"
	"github.com/rgfaber/go-vesca/greenhouse/features/initialize/domain"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/sdk/nats/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewResponder(t *testing.T) {
	// Given
	cfg := config.NewNatsConfig()
	natsBus := NewNatsBus(cfg)
	store := sdk.NewStore()
	decBus := sdk.NewDECBus()
	// And
	aggregate := domain.NewAggregate(store, decBus)
	// When
	rsp := NewResponder(natsBus, aggregate)
	// Then
	assert.NotNil(t, rsp)
	assert.NotNil(t, rsp.natsBus)
	assert.NotNil(t, rsp.aggregate)
}

func TestResponder_ImplementsIResponder(t *testing.T) {
	// Given
	r := NewResponder(nil, nil)
	// When
	ok := sdk.ImplementsIResponder(r)
	// Then
	assert.True(t, ok)
}

func TestResponder_Respond(t *testing.T) {
	// Given
	cfg := config.NewNatsConfig()
	natsBus := NewNatsBus(cfg)
	store := sdk.NewStore()
	decBus := sdk.NewDECBus()
	aggregate := domain.NewAggregate(store, decBus)
	rsp := NewResponder(natsBus, aggregate)
	// And
	id := model.BogusGreenhouseID.Value
	name := model.BogusGreenhouse.Details.Name
	hope := contract.NewHope(id, name)
	// When
	//	fbk, err := rsp.Respond(hope)

}
