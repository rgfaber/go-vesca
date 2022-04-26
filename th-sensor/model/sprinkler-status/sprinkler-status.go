package sprinkler_status

// SprinklerStatus
/////////////////////////////////////////////////////////////////////
type SprinklerStatus int

const (
	Deactivated SprinklerStatus = 0 << iota
	Activated   SprinklerStatus = 1
)
