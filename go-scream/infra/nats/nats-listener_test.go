package nats

import (
	"github.com/rgfaber/go-vesca/sdk/bogus"
	"github.com/rgfaber/go-vesca/sdk/contract"
	"github.com/rgfaber/go-vesca/sdk/dec"
	sdk_nats_config "github.com/rgfaber/go-vesca/sdk/infra/nats/config"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewListener(t *testing.T) {
	// Given
	cfg := sdk_nats_config.NewNatsConfig()
	natsBus := SingletonNatsBus(cfg)

	myFactDeserializer := contract.NewFactDeserializer(bogus.MyFactFromJson)
	myFactHandler := dec.NewFactHandler(myFactDeserializer, func(fact interfaces.IFact) {})
	// When
	lst := NewNatsListener(natsBus, bogus.BOGUS_FACT_TOPIC, myFactHandler)
	// Then
	assert.NotNil(t, lst)
}

func TestListenerImplementsIListener(t *testing.T) {
	// Given
	lst := NewNatsListener(nil, bogus.BOGUS_FACT_TOPIC, nil)
	// When
	ok := interfaces.ImplementsIFactListener(lst)
	// Then
	assert.True(t, ok)

}
