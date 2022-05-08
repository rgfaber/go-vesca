package interfaces

type IIdentity interface {
	Id() string
}

func ImplementsIIdentity(id IIdentity) bool {
	return true
}
