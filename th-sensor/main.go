package main

import (
	"fmt"
	"github.com/rgfaber/go-vesca/th-sensor/domain"
	"os"
)

func main() {
	sensorId := os.Getenv(TH_SENSOR_ID)
	act := domain.NewActor(sensorId)
	if act == nil {
		err := fmt.Errorf("Could not create actor.")
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
