package initialize

type Commands chan *Cmd

func NewCommands(size int) Commands {
	return make(Commands, size)
}
