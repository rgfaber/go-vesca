package event_store

import "github.com/rgfaber/go-vesca/sdk/domain"

type IPlayer interface {
	Replay(aggregateId string) ([]domain.IEvt, error)
}

func ImplementsIPlayer(reader IPlayer) bool {
	return true
}

type PlayerBase struct {
}
