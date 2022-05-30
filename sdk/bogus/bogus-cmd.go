package bogus

import (
	"github.com/rgfaber/go-vesca/sdk/core"
)

type BogusCmd struct {
	ID        core.IIdentity
	SessionId string
	Payload   *BogusPayload
}

func (c *BogusCmd) AggregateId() core.IIdentity {
	return c.ID
}

func (c *BogusCmd) TraceId() string {
	return c.SessionId
}

func NewBogusCmd(id core.IIdentity, sessionId string, payload *BogusPayload) *BogusCmd {
	return &BogusCmd{
		ID:        id,
		SessionId: sessionId,
		Payload:   payload,
	}
}
