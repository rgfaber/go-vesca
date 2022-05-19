package bogus

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

type BogusCmd struct {
	ID        interfaces.IIdentity
	SessionId string
	Payload   *BogusPayload
}

func (c *BogusCmd) AggregateId() interfaces.IIdentity {
	return c.ID
}

func (c *BogusCmd) TraceId() string {
	return c.SessionId
}

func NewBogusCmd(id interfaces.IIdentity, sessionId string, payload *BogusPayload) *BogusCmd {
	return &BogusCmd{
		ID:        id,
		SessionId: sessionId,
		Payload:   payload,
	}
}
