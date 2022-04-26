package model

// SensorStatus
//////////////////////////////////////////////////////////////
type SensorStatus int

const (
	Unknown     SensorStatus = 0 << iota
	Error       SensorStatus = 1
	Created     SensorStatus = 2
	Initialized SensorStatus = 4
	Measuring   SensorStatus = 8
)

func (s SensorStatus) HasFlag(flag SensorStatus) bool {
	return s|flag == s
}

func (s SensorStatus) Unset(flag SensorStatus) SensorStatus {
	return s &^ flag
}

func (s SensorStatus) Set(flag SensorStatus) SensorStatus {
	return s | flag
}
