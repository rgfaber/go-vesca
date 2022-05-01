package sdk

type IIdentity interface {
	Id() string
	Value() string
	Prefix() string
}

func ImplementsIIdentity(id IIdentity) bool {
	return true
}

type ITopic interface {
	Topic() string
}
