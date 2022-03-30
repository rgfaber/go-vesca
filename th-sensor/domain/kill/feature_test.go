package kill

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewFeature(t *testing.T) {
	err := os.Setenv(config.GO_VESCA_TH_SENSOR_ID, sdk.TEST_UUID)
	assert.Nil(t, err)
	cfg := config.NewConfig()
	state := model.NewRoot(cfg)
	assert.NotNil(t, state, "model.Root could not be created for", cfg.SensorId)
	f := NewFeature(state)
	assert.NotNil(t, f, "Feature domain.Kill could not be created.")
}
