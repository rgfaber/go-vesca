package model

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGreenhouse(t *testing.T) {
	cfg := config.NewConfig()
	id := cfg.GreenhouseId()
	gh := NewGreenhouse(id)
	ID := sdk.NewIdentityFrom(config.GO_VESCA_GREENHOUSE_PREFIX, id)
	assert.NotNil(t, gh)
	assert.Equal(t, ID.Value, gh.ID.Value)
}
