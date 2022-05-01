package initialize

type Fact struct {
	aggregateId string
	traceId     string
	payload     interface{}
}

func (f Fact) AggregateId() string {
	return f.aggregateId
}

func (f Fact) TraceId() string {
	return f.traceId
}

func (f Fact) Payload() interface{} {
	return f.payload
}

func NewFact(aggregateId string, traceId string) Fact {
	return Fact{
		aggregateId: aggregateId,
		traceId:     traceId,
		payload:     nil,
	}
}
