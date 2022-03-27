package domain

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"os"
	"strings"
	"testing"
)

func TestNewActor(t *testing.T) {
	err := os.Setenv(model.TH_SENSOR_ID, sdk.TEST_UUID)
	if err != nil {
		t.Errorf("Error setting Environment Variable [%+v]", model.TH_SENSOR_ID)
	}
	model.SensorId = os.Getenv(model.TH_SENSOR_ID)
	act := NewActor(sdk.TEST_UUID)
	if act == nil {
		t.Errorf("Actor could not be created")
	}
	if act.Id == nil {
		t.Errorf("Actor has no Identity")
	}
	if act.Id.Prefix != model.TH_SENSOR {
		t.Errorf("Actor has wrong Prefix. Expected [%+v] Got [%+v]", model.TH_SENSOR, act.Id.Prefix)
	}
	ts := strings.Replace(model.SensorId, "-", "", -1)
	if act.Id.Value != ts {
		t.Errorf("Actor has wrong Value. Expected [%+v] Got [%+v]", ts, act.Id.Prefix)
	}
}
