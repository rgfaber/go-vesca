package measure

import (
	"github.com/rgfaber/go-vesca/sdk"
)

const CMD_TOPIC = "th_sensor:measure"

type Cmd struct{}

func (c *Cmd) AggregateId() sdk.IIdentity {
	//TODO implement me
	panic("implement me")
}

func NewCmd() *Cmd {
	return &Cmd{}
}
