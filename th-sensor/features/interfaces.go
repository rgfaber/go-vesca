package features

type IEmitter interface {
	Emit(evt interface{})
}

type IResponder interface {
	Respond(cmd interface{})
}

type IFeature interface {
	Run()
	Listen()
	Respond()
}
