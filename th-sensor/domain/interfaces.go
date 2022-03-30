package domain

type ISupervisor interface {
	Supervise()
}

type IFeature interface {
	Run()
}
