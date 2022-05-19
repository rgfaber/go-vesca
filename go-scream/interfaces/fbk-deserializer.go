package interfaces

type IFbkDeserializer interface {
	Deserialize([]byte) IFbk
}

func ImplementsIFbkDeserializer(deserializer IFbkDeserializer) bool {
	return true
}
