package interfaces

type ICmd interface {
	AggregateId() IIdentity
	TraceId() string
}

func ImplementsICmd(cmd ICmd) bool {
	return true
}
