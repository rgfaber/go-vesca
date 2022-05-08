package interfaces

type IFactDeserializer interface {
	Deserialize([]byte) IFact
}

func ImplementsIFactDeserializer(fd IFactDeserializer) bool {
	return true
}
