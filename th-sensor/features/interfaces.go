package domain

type IFeature interface {
	Run()
	Listen()
	Respond()
}

type IEmitter interface {
	Emit(evt interface{})
}

type IResponder interface {
	Respond(cmd interface{})
}
