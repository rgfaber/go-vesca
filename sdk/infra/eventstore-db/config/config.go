package config

import (
	"github.com/rgfaber/go-vesca/sdk/infra/eventstore-db/envars"
	"os"
)

const (
	DefaultConnectionString = "esdb://127.0.0.1:2113?tls=false&keepAliveTimeout=10000&keepAliveInterval=10000"
)

type IConfig interface {
	ConnectionString() string
}

type Config struct {
}

func (c *Config) ConnectionString() string {
	str := os.Getenv(envars.ESDB_CONNECTION_STRING)
	if str == "" {
		return DefaultConnectionString
	}
	return str
}

func NewConfig() IConfig {
	return &Config{}
}
