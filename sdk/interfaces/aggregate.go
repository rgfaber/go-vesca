package interfaces

type IAggregate interface {
	Attempt(cmd ICmd) IFbk
	Apply(evt IEvt)
}

func ImplementsIAggregate(aggregate IAggregate) bool {
	return true
}
