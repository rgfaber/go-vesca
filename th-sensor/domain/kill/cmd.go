package kill

const CMD_TOPIC = "th_sensor:kill"

type Cmd struct {
	SessionId string `json:"session_id"`
}

func NewCnd(sessionId string) *Cmd {
	return &Cmd{
		SessionId: sessionId,
	}
}

func (c *Cmd) ToEvt() *Evt {
	return &Evt{SessionId: c.SessionId}
}
