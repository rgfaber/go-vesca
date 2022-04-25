package initialize

import (
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRsp(t *testing.T) {
	id := model.NewTHSensorTestID()
	r := NewRsp(id, model.Initialized, true, "test")
	assert.NotNil(t, r)
}
