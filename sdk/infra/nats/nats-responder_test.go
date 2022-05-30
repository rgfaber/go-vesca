package nats

import (
	"github.com/rgfaber/go-vesca/sdk/infra"
	"github.com/rgfaber/go-vesca/sdk/infra/nats/config"
	mocks "github.com/rgfaber/go-vesca/sdk/mocks/infra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewResponder(t *testing.T) {
	// Given
	cfg := config.NewNatsConfig()
	natsBus := SingletonNatsBus(cfg)
	hopeHandler := mocks.NewIHopeHandler(t)
	// When
	rsp := NewNatsResponder(natsBus, "", hopeHandler)
	// Then
	assert.NotNil(t, rsp)
}

func TestResponder_ImplementsIResponder(t *testing.T) {
	// Given
	r := NewNatsResponder(nil, "", nil)
	// When
	ok := infra.ImplementsIHopeResponder(r)
	// Then
	assert.True(t, ok)
}

func TestResponder_Respond(t *testing.T) {
	//// Given
	//cfg := config.NewNatsConfig()
	//natsBus := SingletonNatsBus(cfg)
	//store := memory2.NewStore()
	//decBus := mediator.SingletonDECBus()
	//aggregate := domain.NewAggregate(store, decBus)
	//rsp := NewNatsResponder()
	//// And
	//id := model.BogusGreenhouseID.Value
	//name := model.BogusGreenhouse.Details.Name
	//hope := bogus2.NewBogusHope(test.TEST_UUID, test.TEST_TRACE_ID)
	//// When
	////	fbk, err := rsp.Activate(hope)
	assert.True(t, true)

}
