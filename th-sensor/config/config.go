package config

import (
	"github.com/nats-io/nats.go"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/envars"
	"os"
)

const (
	GO_VESCA_TH_SENSOR_PREFIX  = "vesca_th_sensor"
	GO_VESCA_GREENHOUSE_PREFIX = "vesca_greenhouse"
)

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) NATSUrl() string {
	url := os.Getenv(envars.GO_VESCA_NATS_URL)
	if url == "" {
		return nats.DefaultURL
	}
	return url
}

func (c *Config) NATSSUser() string {
	usr := os.Getenv(envars.GO_VESCA_NATS_USER)
	if usr == "" {
		return "a"
	}
	return usr
}

func (c *Config) NATSPwd() string {
	pwd := os.Getenv(envars.GO_VESCA_NATS_PWD)
	if pwd == "" {
		return "a"
	}
	return pwd
}

func (c *Config) SensorId() string {
	id := os.Getenv(envars.GO_VESCA_TH_SENSOR_ID)
	if id == "" {
		return sdk.TEST_UUID
	}
	return id
}

func (c *Config) SensorName() string {
	name := os.Getenv(envars.GO_VESCA_TH_SENSOR_NAME)
	if name == "" {
		return "Sensor 0"
	}
	return name
}

func (c *Config) GreenhouseName() string {
	name := os.Getenv(envars.GO_VESCA_GREENHOUSE_NAME)
	if name == "" {
		return "Greenhouse 0 - Tomatoes"
	}
	return name
}

func (c *Config) GreenhouseId() string {
	id := os.Getenv(envars.GO_VESCA_GREENHOUSE_ID)
	if id == "" {
		return sdk.TEST_UUID
	}
	return id

}
