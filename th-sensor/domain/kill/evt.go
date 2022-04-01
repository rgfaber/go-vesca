package kill

const EVT_TOPIC = "th_sensor:killed"

type Evt struct {
	SessionId string
}

func NewEvt() *Evt {
	return &Evt{}
}
