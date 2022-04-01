package measure

import (
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEvt(t *testing.T) {
	m := model.NewMeasurement(23, 40)
	evt := NewEvt(m)
	assert.NotNil(t, evt)
}
