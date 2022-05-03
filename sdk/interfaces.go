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
	TraceId() string
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
	Serialize(hope IHope) string
	Deserialize(s string) IHope
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

type IEmitter interface {
	Emit(fact IFact)
	Bus() IDECBus
	Topic() string
	Serialize(fact IFact) []byte
	Deserialize(s []byte) IFact
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

type IIdentity interface {
	Id() string
	Value() string
	Prefix() string
}

func ImplementsIIdentity(id IIdentity) bool {
	return true
}

type ITopic interface {
	Topic() string
}

func ImplementsITopic(t ITopic) bool {
	return true
}
