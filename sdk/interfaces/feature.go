package interfaces

type IFeature interface {
	Run()
}

func ImplementsIFeature(feature IFeature) bool {
	return true
}
