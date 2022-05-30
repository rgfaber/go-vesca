package config

import (
	"github.com/rgfaber/go-vesca/greenhouse/config/envars"
	"github.com/rgfaber/go-vesca/sdk/core/mocks"
	"os"
)

const (
	GO_VESCA_TH_SENSOR_PREFIX  = "thsensor"
	GO_VESCA_GREENHOUSE_PREFIX = "greenhouse"
	GO_VESCA_FAN_PREFIX        = "fan"
	GO_VESCA_SPRINKLER_PREFIX  = "sprinkler"
)

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) SensorId() string {
	id := os.Getenv(envars.GO_VESCA_TH_SENSOR_ID)
	if id == "" {
		return mocks.TEST_UUID
	}
	return id
}

func (c *Config) SensorName() string {
	return os.Getenv(envars.GO_VESCA_TH_SENSOR_NAME)
}

func (c *Config) GreenhouseName() string {
	return os.Getenv(envars.GO_VESCA_GREENHOUSE_NAME)
}

func (c *Config) GreenhouseId() string {
	id := os.Getenv(envars.GO_VESCA_GREENHOUSE_ID)
	if id == "" {
		return mocks.TEST_UUID
	}
	return id
}

func (c *Config) SprinklerId() string {
	id := os.Getenv(envars.GO_VESCA_SPRINKLER_ID)
	if id == "" {
		return mocks.TEST_UUID
	}
	return id
}

func (c *Config) SprinklerName() string {
	return os.Getenv(envars.GO_VESCA_SPRINKLER_NAME)
}

func (c *Config) FanId() string {
	id := os.Getenv(envars.GO_VESCA_FAN_ID)
	if id == "" {
		return mocks.TEST_UUID
	}
	return id
}

func (c *Config) FanName() string {
	return os.Getenv(envars.GO_VESCA_FAN_NAME)
}
