package kill

type Evt struct {
	SessionId string
}

func NewEvt(sessionId string) *Evt {
	return &Evt{SessionId: sessionId}
}
