package nats

import (
	"github.com/rgfaber/go-vesca/sdk/bogus"
	"github.com/rgfaber/go-vesca/sdk/domain"
	"github.com/rgfaber/go-vesca/sdk/infra"
	sdk_nats_config "github.com/rgfaber/go-vesca/sdk/infra/nats/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewListener(t *testing.T) {
	// Given
	cfg := sdk_nats_config.NewNatsConfig()
	natsBus := SingletonNatsBus(cfg)

	myFactHandler := infra.NewFactHandler(func(fact domain.IFact) {}, func(data []byte) domain.IFact { return nil })
	// When
	lst := NewNatsListener(natsBus, bogus.BOGUS_FACT_TOPIC, myFactHandler)
	// Then
	assert.NotNil(t, lst)
}

func TestListenerImplementsIListener(t *testing.T) {
	// Given
	lst := NewNatsListener(nil, bogus.BOGUS_FACT_TOPIC, nil)
	// When
	ok := infra.ImplementsIFactListener(lst)
	// Then
	assert.True(t, ok)

}
