package initialize

type Events chan *Evt

func NewEvents(size int) Events {
	return make(Events, size)
}
