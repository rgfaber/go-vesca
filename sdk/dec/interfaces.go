package dec

import (
	"github.com/rgfaber/go-vesca/sdk"
)

type IStore interface {
	Load(id string) interface{}
	Save(model interface{})
}

func ImplemementsIStore(store IStore) bool {
	return true
}

type ICmd interface {
	AggregateId() sdk.IIdentity
}

func ImplementsICmd(cmd ICmd) bool {
	return true
}

type IEvt interface {
	AggregateId() sdk.IIdentity
}

func ImplementsIEvt(evt IEvt) bool {
	return true
}

type IFbk interface {
	AggregateId() string
	Status() int
	IsSuccess() bool
	TraceId() string
}

func ImplementsIFbk(fbk IFbk) bool {
	return true
}

type IAggregate interface {
	Attempt(cmd ICmd) (IFbk, error)
	Apply(evt IEvt)
}

func ImplementsIAggregate(aggregate IAggregate) bool {
	return true
}

type IFact interface {
	AggregateId() string
	TraceId() string
	Payload() interface{}
}

func ImplementsIFact(fact IFact) bool {
	return true
}

type IHope interface {
	AggregateId() string
	TraceId() string
	Payload() interface{}
}

func ImplementsIHope(hope IHope) bool {
	return true
}

type IFeature interface {
	Bus() IDECBus
	Store() IStore
	Aggregate() IAggregate
	Run()
	Listen()
	Respond()
}

func ImplemmentsIFeature(feature IFeature) bool {
	return true
}

type IEmitter interface {
	Emit(fact IFact)
	Bus() IDECBus
}

func ImplementsIEmitter(emitter IEmitter) bool {
	return true
}

type IResponder interface {
	Respond(hope IHope) (IFbk, error)
	Bus() IDECBus
}

func ImplementsIResponder(rsp IResponder) bool {
	return true
}
