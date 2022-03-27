package initialize

type Handler struct {
	Commands chan Cmd
	Events   chan Evt
}

func NewHandler(cmds chan Cmd, evts chan Evt) *Handler {
	return &Handler{
		Commands: make(chan Cmd),
		Events:   make(chan Evt),
	}
}
