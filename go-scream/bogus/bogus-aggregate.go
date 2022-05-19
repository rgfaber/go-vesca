package bogus

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

type BogusAggregate struct {
}

func (a *BogusAggregate) Attempt(cmd interfaces.ICmd) interfaces.IFbk {
	return nil
}

func (a *BogusAggregate) Raise(evt interfaces.IEvt) {

}

func (a *BogusAggregate) Apply(evt interfaces.IEvt) {
}

func NewBogusAggregate() *BogusAggregate {
	return &BogusAggregate{}
}
