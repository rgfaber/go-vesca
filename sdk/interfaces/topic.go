package interfaces

type ITopic interface {
	Topic() string
}

func ImplementsITopic(topic ITopic) bool {
	return true
}
