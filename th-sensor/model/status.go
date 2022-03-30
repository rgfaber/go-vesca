package model

type SensorStatus int

const (
	Unknown     SensorStatus = 0 << iota
	Error       SensorStatus = 1
	Created     SensorStatus = 2
	Initialized SensorStatus = 4
	Killed      SensorStatus = 8
)

func (this SensorStatus) HasFlag(flag SensorStatus) bool {
	return this|flag == this
}

func (this SensorStatus) Unset(flag SensorStatus) SensorStatus {
	return this &^ flag
}

func (this SensorStatus) Set(flag SensorStatus) SensorStatus {
	return this | flag
}
