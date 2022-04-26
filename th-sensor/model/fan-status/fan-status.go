package fan_status

// FanStatus
//////////////////////////////////////////////////////////////////////
type FanStatus int

const (
	Deactivated FanStatus = 0 << iota
	Activated   FanStatus = 1
)
