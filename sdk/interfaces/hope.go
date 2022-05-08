package interfaces

type IHope interface {
	AggregateId() string
	TraceId() string
}

func ImplementsIHope(hope IHope) bool {
	return true
}
