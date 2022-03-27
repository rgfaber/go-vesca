package main

import (
	"fmt"
	"github.com/rgfaber/go-vesca/th-sensor/domain"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"os"
)

func main() {
	err := os.Setenv(model.TH_SENSOR_ID, "fc6e30ed-2194-42b8-9d80-1ac8718e9655")
	if err != nil {
		panic(err)
	}
	sensorId := os.Getenv(model.TH_SENSOR_ID)
	act := domain.NewActor(sensorId)
	if act == nil {
		err := fmt.Errorf("Could not create actor.")
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	act.Run()
	select {}
}
