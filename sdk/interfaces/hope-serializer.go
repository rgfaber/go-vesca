package interfaces

type IHopeSerializer interface {
	Serialize(fact IHope) ([]byte, error)
}

func ImplementsIHopeSerializer(serializer IHopeSerializer) bool {
	return true
}
