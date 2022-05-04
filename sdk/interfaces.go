package sdk

type IStore interface {
	Load(id string) interface{}
	Save(model interface{})
}

func ImplemementsIStore(store IStore) bool {
	return true
}

type ICmd interface {
	AggregateId() IIdentity
}

func ImplementsICmd(cmd ICmd) bool {
	return true
}

type IEvt interface {
	AggregateId() IIdentity
}

func ImplementsIEvt(evt IEvt) bool {
	return true
}

type IFbk interface{}

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
	Topic()
}

func ImplementsIFact(fact IFact) bool {
	return true
}

type IHope interface {
	Topic() string
}

func ImplementsIHope(hope IHope) bool {
	return true
}

type IFeature interface {
	Bus() IDECBus
	Store() IStore
	Aggregate() IAggregate
	Run()
	StartListening()
	StartResponding()
}

func ImplemmentsIFeature(feature IFeature) bool {
	return true
}

type IIdentity interface {
	Id() string
}

func ImplementsIIdentity(id IIdentity) bool {
	return true
}

type IListener interface {
	Listen(func(fact IFact))
}

func ImplementsIListener(ls IListener) bool {
	return true
}

type IResponder interface {
	Respond(IHope) (IFbk, error)
}

func ImplementsIResponder(rsp IResponder) bool {
	return true
}

type IEmitter interface {
	Emit(fact IFact)
}

func ImplementsIEmitter(emitter IEmitter) bool {
	return true
}

type IRequester interface {
	Request(IHope) (IFbk, error)
}

func ImplementsIRequester(requester IRequester) bool {
	return true
}
