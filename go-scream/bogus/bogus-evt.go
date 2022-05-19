package bogus

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

type Evt struct {
	ID        interfaces.IIdentity
	SessionId string
	Payload   *BogusPayload
}

func NewEvt(id interfaces.IIdentity, sessionId string, payload *BogusPayload) *Evt {
	return &Evt{
		ID:        id,
		SessionId: sessionId,
		Payload:   payload,
	}
}
