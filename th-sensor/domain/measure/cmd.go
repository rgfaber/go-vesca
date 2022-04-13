package measure

const CMD_TOPIC = "th_sensor:measure"

type Cmd struct{}

func NewCmd() *Cmd {
	return &Cmd{}
}
