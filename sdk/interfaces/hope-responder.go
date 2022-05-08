package interfaces

type IHopeResponder interface {
	Respond([]byte) []byte
}

func ImplementsIHopeResponder(rsp IHopeResponder) bool {
	return true
}
