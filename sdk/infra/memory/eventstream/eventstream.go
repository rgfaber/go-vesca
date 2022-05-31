package eventstream

import (
	"github.com/rgfaber/go-vesca/sdk/core"
	"github.com/rgfaber/go-vesca/sdk/domain"
	event_store "github.com/rgfaber/go-vesca/sdk/infra/memory/event-store"
	memory_mediator "github.com/rgfaber/go-vesca/sdk/infra/memory/mediator"
	"sort"
)

const (
	AGGREGATE_TOPIC = "aggregate"
)

type EventStream struct {
	Mediator memory_mediator.IDECBus
	Events   map[string]domain.IEvt
	Player   event_store.IPlayer
}

func (m EventStream) AppendEvent(aggregateId core.IIdentity) error {
	//TODO implement me
	panic("implement me")
}

func (m EventStream) Broadcast(events []domain.IEvt) error {
	//TODO implement me
	panic("implement me")
}

func (m EventStream) LoadEvents(aggregateId core.IIdentity) error {
	events, err := m.Player.Replay(aggregateId.Id())
	sort.SliceStable(events, func(i, j int) bool {
		return events[i].Order() < events[j].Order()
	})
	if err == nil {
		return err
	}
	for _, evt := range events {
		m.Mediator.Publish(AGGREGATE_TOPIC, evt)
	}
	return nil
}

func NewEventStream(mediator memory_mediator.IDECBus, reader event_store.IPlayer) domain.IEventStream {
	return &EventStream{
		Events:   make(map[string]domain.IEvt),
		Mediator: mediator,
		Player:   reader,
	}
}
