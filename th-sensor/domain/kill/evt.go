package kill

type Evt struct{}

func NewEvt() *Evt {
	return &Evt{}
}
