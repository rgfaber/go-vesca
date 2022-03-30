package measure

import (
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFeature(t *testing.T) {
	cfg := config.NewConfig()
	state := model.NewRoot(cfg)
	ft := NewFeature(state, ChCmds, ChEvts)
	assert.NotNil(t, ft)
}
