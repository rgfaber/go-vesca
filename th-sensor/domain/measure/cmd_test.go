package measure

import (
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCmd(t *testing.T) {
	m := model.NewMeasurement(16, 25)
	cmd := NewCmd(m)
	assert.NotNil(t, cmd)
}
