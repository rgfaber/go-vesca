package sdk

type IIdentity interface {
	Id() (string string)
}

type ITopic interface {
	Topic() string
}
