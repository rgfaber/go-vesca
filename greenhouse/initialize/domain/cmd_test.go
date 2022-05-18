package domain

import (
	bogus_model "github.com/rgfaber/go-vesca/greenhouse/model"
	testing2 "github.com/rgfaber/go-vesca/sdk/core/test"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCmd(t *testing.T) {
	id := bogus_model.BogusGreenhouseID
	c := NewCmd(id, testing2.TEST_TRACE_ID,
		bogus_model.NewDetails("greenhouse", ""),
		bogus_model.BogusSettings,
		bogus_model.BogusSensor,
		bogus_model.BogusFan,
		bogus_model.BogusSprinkler)
	assert.NotNil(t, c)
	assert.Equal(t, bogus_model.BogusTemp, c.Settings.Temperature)
	assert.Equal(t, bogus_model.BogusHumidity, c.Settings.Humidity)
}

func TestCmd_ImplementsICmd(t *testing.T) {
	// Given
	cmd := NewCmd(nil, "", nil, nil, nil, nil, nil)
	// When
	ok := interfaces.ImplementsICmd(cmd)
	// Then
	assert.True(t, ok)
}
