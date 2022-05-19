package interfaces

type IFbkSerializer interface {
	Serialize(fbk IFbk) ([]byte, error)
}

func ImplementsIFbkSerializer(serializer IFbkSerializer) bool {
	return true
}
