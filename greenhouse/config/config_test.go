package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfig(t *testing.T) {
	cfg := NewConfig()
	assert.NotNil(t, cfg, "Could not create infra.Config")
	fmt.Println("Config |> GreenhouseId", cfg.GreenhouseId())
	fmt.Println("Config |> GreenhouseName", cfg.GreenhouseName())
}
