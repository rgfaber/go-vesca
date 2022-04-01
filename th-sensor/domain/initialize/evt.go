package initialize

const EVT_TOPIC = "th_sensor:initialized"

type Evt struct{}

func NewEvt() *Evt {
	return &Evt{}
}
