package kill

type IEmitter interface {
	Emit(evt Evt)
}

type IResponder interface {
	Respond(cmd Cmd)
}
