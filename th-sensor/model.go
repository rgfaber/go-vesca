package main

type thMeasurement struct {
	sensorId    string  `json:"sensor_id"`
	temperature float64 `json:"temperature"`
	humidity    float64 `json:"humidity"`
}
