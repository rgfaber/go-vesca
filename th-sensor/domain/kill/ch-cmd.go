package kill

type ChCmd chan *Cmd

func NewChCmd(size int) ChCmd {
	return make(ChCmd, size)
}
