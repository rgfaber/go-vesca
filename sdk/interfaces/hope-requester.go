package interfaces

type IHopeRequester interface {
	Request(IHope) (IFbk, error)
}

func ImplementsIRequester(requester IHopeRequester) bool {
	return true
}
