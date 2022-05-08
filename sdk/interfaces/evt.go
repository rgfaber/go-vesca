package interfaces

type IEvt interface {
	AggregateId() IIdentity
	TraceId() string
}

func ImplementsIEvt(evt IEvt) bool {
	return true
}
