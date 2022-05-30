package bogus

import (
	"github.com/rgfaber/go-vesca/sdk/core"
	"github.com/rgfaber/go-vesca/sdk/domain"
)

type BogusAggregate struct {
}

func (a *BogusAggregate) LoadEvents(aggregateId core.IIdentity) []domain.IEvt {
	//TODO implement me
	panic("implement me")
}

func (a *BogusAggregate) AppendEvent(evt domain.IEvt) {
	//TODO implement me
	panic("implement me")
}

func (a *BogusAggregate) Attempt(cmd domain.ICmd) domain.IFbk {
	return nil
}

func (a *BogusAggregate) Raise(evt domain.IEvt) {

}

func (a *BogusAggregate) Apply(evt domain.IEvt) {
}

func NewBogusAggregate() *BogusAggregate {
	return &BogusAggregate{}
}
