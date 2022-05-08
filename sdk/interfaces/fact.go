package interfaces

type IFact interface {
	AggregateId() string
	TraceId() string
}

func ImplementsIFact(fact IFact) bool {
	return true
}
