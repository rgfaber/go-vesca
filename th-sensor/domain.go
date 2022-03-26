package main

type sensor struct {
	id string `json:"id"`
}

type sensorCreated struct {
	sensor sensor
}

func NewSensor(id string) (d *sensor) {
	return &sensor{
		id: id,
	}
}
