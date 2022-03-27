package model

type SensorStatus int

const (
	Unknown SensorStatus = iota
	Error
	Created
	Initialized
	Activated
	Deactivated
)
