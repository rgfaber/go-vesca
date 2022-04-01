package initialize

const CMD_TOPIC = "th_sensor:initialize"

type Cmd struct{}

func NewCmd() *Cmd {
	return &Cmd{}
}
