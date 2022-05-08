package measure

import (
	"github.com/rgfaber/go-vesca/greenhouse/features/measure/domain"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFeature(t *testing.T) {
	ft := NewFeature(sdk.NewDECBus(), sdk.NewStore())
	assert.NotNil(t, ft)
}

func TestMeasure(t *testing.T) {
	ft := NewFeature(sdk.NewDECBus(), sdk.NewStore())
	assert.NotNil(t, ft)
	go ft.Respond()
	ft.bus.Publish(domain.CMD_TOPIC, domain.NewCmd())
}
