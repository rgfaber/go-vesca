package domain

import (
	bogus "github.com/rgfaber/go-vesca/greenhouse/model"
	bogus_model "github.com/rgfaber/go-vesca/greenhouse/model"
	testing2 "github.com/rgfaber/go-vesca/sdk/core/mocks"
	"github.com/rgfaber/go-vesca/sdk/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCmd(t *testing.T) {
	id := bogus.BogusGreenhouseID
	c := NewCmd(id, testing2.TEST_TRACE_ID,
		bogus_model.NewDetails("greenhouse", ""),
		bogus.BogusSettings,
		bogus.BogusSensor,
		bogus.BogusFan,
		bogus.BogusSprinkler)
	assert.NotNil(t, c)
	assert.Equal(t, bogus.BogusTemp, c.Settings.Temperature)
	assert.Equal(t, bogus.BogusHumidity, c.Settings.Humidity)
}

func TestCmd_ImplementsICmd(t *testing.T) {
	// Given
	cmd := NewCmd(nil, "", nil, nil, nil, nil, nil)
	// When
	ok := domain.ImplementsICmd(cmd)
	// Then
	assert.True(t, ok)
}
