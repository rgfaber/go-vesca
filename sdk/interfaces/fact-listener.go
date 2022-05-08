package interfaces

type IFactListener interface {
	Listen(func(fact IFact))
}

func ImplementsIListener(ls IFactListener) bool {
	return true
}
