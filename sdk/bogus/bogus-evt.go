package bogus

import (
	"github.com/rgfaber/go-vesca/sdk/core"
)

type Evt struct {
	ID        core.IIdentity
	SessionId string
	Payload   *BogusPayload
}

func NewEvt(id core.IIdentity, sessionId string, payload *BogusPayload) *Evt {
	return &Evt{
		ID:        id,
		SessionId: sessionId,
		Payload:   payload,
	}
}
