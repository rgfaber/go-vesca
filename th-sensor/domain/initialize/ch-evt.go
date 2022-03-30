package initialize

type ChEvt chan *Evt

func NewChEvt(size int) ChEvt {
	return make(ChEvt, size)
}
