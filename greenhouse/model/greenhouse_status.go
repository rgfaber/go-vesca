package model

// GreenhouseStatus
//////////////////////////////////////////////////////////////
type GreenhouseStatus int

const (
	Unknown            GreenhouseStatus = 0 << iota
	Error              GreenhouseStatus = 1
	Created            GreenhouseStatus = 2
	Initialized        GreenhouseStatus = 4
	FanActivated       GreenhouseStatus = 8
	SprinklerActivated GreenhouseStatus = 16
	FanReplaced        GreenhouseStatus = 32
	SensorReplaced     GreenhouseStatus = 64
	SprinklerReplaced  GreenhouseStatus = 128
)

func (s GreenhouseStatus) HasFlag(flag GreenhouseStatus) bool {
	return s|flag == s
}

func (s GreenhouseStatus) Unset(flag GreenhouseStatus) GreenhouseStatus {
	return s &^ flag
}

func (s GreenhouseStatus) Set(flag GreenhouseStatus) GreenhouseStatus {
	return s | flag
}
