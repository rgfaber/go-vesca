package config

import (
	"github.com/nats-io/nats.go"
	"github.com/rgfaber/go-vesca/sdk/nats/envars"
	"os"
)

type NatsConfig struct {
}

func (c *NatsConfig) NATSUrl() string {
	url := os.Getenv(envars.NATS_URL)
	if url == "" {
		return nats.DefaultURL
	}
	return url
}

func (c *NatsConfig) NATSUser() string {
	usr := os.Getenv(envars.NATS_USER)
	if usr == "" {
		return "a"
	}
	return usr
}

func (c *NatsConfig) NATSPwd() string {
	pwd := os.Getenv(envars.NATS_PWD)
	if pwd == "" {
		return "a"
	}
	return pwd
}

type INatsConfig interface {
	NATSUrl() string
	NATSUser() string
	NATSPwd() string
}

func ImplementsINatsConfig(cfg INatsConfig) bool {
	return true
}

func NewNatsConfig() INatsConfig {
	return &NatsConfig{}
}
