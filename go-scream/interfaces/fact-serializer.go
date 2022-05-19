package interfaces

type IFactSerializer interface {
	Serialize(fact IFact) ([]byte, error)
}

func ImplementsIFactSerializer(serializer IFactSerializer) bool {
	return true
}
