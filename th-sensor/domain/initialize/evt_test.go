package initialize

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEvt(t *testing.T) {
	m := model.NewMeasurement(15.2, 50)
	aggregateId = sdk.NewIdentity()
	evt := NewEvt(*m)
	assert.NotNil(t, evt)
}
