package contract

type Hope struct {
	aggregateId string
	traceId     string
	Name        string
}

func (h *Hope) AggregateId() string {
	return h.aggregateId
}

func (h *Hope) TraceId() string {
	return h.traceId
}

func NewHope(id string, traceId string, name string) *Hope {
	return &Hope{
		aggregateId: id,
		traceId:     traceId,
		Name:        name,
	}
}
