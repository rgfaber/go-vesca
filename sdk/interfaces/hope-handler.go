package interfaces

type IHopeHandler interface {
	Handle(hope []byte) []byte
}

func ImplementsIHopeHandler(handler IHopeHandler) bool {
	return true
}
