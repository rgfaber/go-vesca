package main

import (
	"fmt"
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"os"
)

func main() {
	cfg := config.NewConfig()
	supervisor := NewSupervisor(cfg, Features)
	if supervisor == nil {
		err := fmt.Errorf("Could not create Supervisor.")
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	supervisor.Supervise()
	select {}
}
