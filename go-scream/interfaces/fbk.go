package interfaces

type IFbk interface {
	AggregateId() string
	TraceId() string
	IsSuccess() bool
	Error() string
}

func ImplementsIFbk(fbk IFbk) bool {
	return true
}
