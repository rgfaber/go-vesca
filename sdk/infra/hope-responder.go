package infra

type IHopeResponder interface {
	Activate()
}

func ImplementsIHopeResponder(rsp IHopeResponder) bool {
	return true
}

func BuildResponders(responders ...IHopeResponder) []IHopeResponder {
	return responders
}
