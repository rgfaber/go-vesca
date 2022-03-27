package kill

import (
	"os"
	"time"
)

type Handler struct {
	Commands chan Cmd
	Events   chan Evt
}

func NewHandler() *Handler {
	return &Handler{
		Commands: make(chan Cmd),
		Events:   make(chan Evt),
	}
}

func (h *Handler) Run() {
	for k := range h.Commands {
		go func(cmd Cmd) {
			time.Sleep(5 * time.Second)
			os.Exit(0)
		}(k)
	}
}
