package interfaces

type IFactHandler interface {
	Handle(data []byte)
}

func ImplementsIFactHandler(handler IFactHandler) bool {
	return true
}
