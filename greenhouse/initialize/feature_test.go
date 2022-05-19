package initialize

import (
	"fmt"
	core_test "github.com/rgfaber/go-vesca/go-scream/core/test"
	"github.com/rgfaber/go-vesca/go-scream/dec"
	"github.com/rgfaber/go-vesca/go-scream/infra/nats"
	"github.com/rgfaber/go-vesca/go-scream/interfaces"
	"github.com/rgfaber/go-vesca/greenhouse/config/envars"
	"github.com/rgfaber/go-vesca/greenhouse/initialize/contract"
	"github.com/rgfaber/go-vesca/greenhouse/initialize/domain"
	"github.com/rgfaber/go-vesca/greenhouse/initialize/infra/memory"
	initialize_nats "github.com/rgfaber/go-vesca/greenhouse/initialize/infra/nats"
	"github.com/rgfaber/go-vesca/greenhouse/initialize/topics"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/stretchr/testify/assert"
	"os"
	"sync"
	"testing"
)

func setEnv() {
	err := os.Setenv(envars.GO_VESCA_TH_SENSOR_ID, core_test.TEST_UUID)
	if err != nil {
		panic(err)
	}
}

func TestNewFeature(t *testing.T) {
	// Given
	setEnv()

	// When
	f := Feature
	// Then
	assert.NotNil(t, f, "Feature domain.Initialize could not be created.")
}

func TestFeature_RespondToInitializeCmd(t *testing.T) {
	// Given
	setEnv()
	receivedEvent := false
	cmd := domain.NewCmd(model.BogusGreenhouseID,
		core_test.TEST_TRACE_ID,
		model.BogusDetails,
		model.BogusSettings,
		model.BogusSensor,
		model.BogusFan,
		model.BogusSprinkler)
	aggregate := domain.Root
	// When
	memory.Mediator.Subscribe(topics.EVT_TOPIC, func(evt interfaces.IEvt) {
		e := evt.(*domain.Evt)
		receivedEvent = e != nil
	})
	// And
	fbk := aggregate.Attempt(cmd)
	// then
	assert.NotNil(t, fbk)
	assert.True(t, receivedEvent)
}

func TestImplementsIFeature(t *testing.T) {
	// Given
	f := Feature
	// When
	imp := interfaces.ImplementsIFeature(f)
	// Then
	assert.True(t, imp)
}

func Test_Fact2Json(t *testing.T) {
	// Given
	fact := contract.NewFact(core_test.TEST_UUID, core_test.TEST_TRACE_ID, model.BogusGreenhouse)
	// When
	json := contract.Fact2Json(fact)
	// Then
	assert.NotEmpty(t, json)
}

func Test_Hope2Json(t *testing.T) {
	// Given
	hope := contract.NewHope(model.BogusGreenhouseID.Id(),
		core_test.TEST_TRACE_ID,
		model.BogusDetails,
		model.BogusSettings,
		model.BogusSensor,
		model.BogusFan,
		model.BogusSprinkler)
	// When
	json := contract.Hope2Json(hope)
	// Then
	assert.NotEmpty(t, json)
}

func Test_Hope2Cmd(t *testing.T) {
	// Given
	hope := contract.NewHope(model.BogusGreenhouseID.Id(),
		core_test.TEST_TRACE_ID,
		model.BogusDetails,
		model.BogusSettings,
		model.BogusSensor,
		model.BogusFan,
		model.BogusSprinkler)
	// When
	cmd := domain.Hope2Cmd(hope)
	c := cmd.(*domain.Cmd)
	// Then
	assert.NotNil(t, cmd)
	assert.Equal(t, hope.AggregateId(), cmd.AggregateId().Id())
	assert.Equal(t, hope.Details, c.Details)
}

var lock = sync.Mutex{}

func Test_Emit2Nats(t *testing.T) {
	// Given
	factHandled := false
	i := 0
	evt := domain.NewEvt(model.BogusGreenhouseID, core_test.TEST_TRACE_ID, model.BogusGreenhouse)
	testHandler := func(fact interfaces.IFact) {
		lock.Lock()
		defer lock.Unlock()
		factHandled = true
		f := fact.(*contract.Fact)
		fmt.Printf("Handled Fact #%+v => [%+v]\n", i, *f)
		i++
	}
	factHandler := dec.NewFactHandler(testHandler, contract.Json2Fact)

	factListener := nats.NewNatsListener(
		initialize_nats.Bus,
		topics.NATS_FACT_TOPIC,
		factHandler)

	factListener.Activate()

	// When
	initialize_nats.Emit2Nats(evt)

	//	time.Sleep(2 * time.Second)

	initialize_nats.Emit2Nats(evt)

	//	time.Sleep(2 * time.Second)

	initialize_nats.Emit2Nats(evt)

	//	time.Sleep(2 * time.Second)

	initialize_nats.Emit2Nats(evt)

	//	time.Sleep(2 * time.Second)
	// Then
	assert.True(t, factHandled)

	assert.Equal(t, 3, i)

}

func Test_Respond2Hope(t *testing.T) {
	// Given
	id := model.BogusGreenhouseID.Id()
	hope := contract.NewHope(id, core_test.TEST_TRACE_ID, model.BogusDetails, model.BogusSettings, model.BogusSensor, model.BogusFan, model.BogusSprinkler)
	hopeHandler = dec.NewHopeHandler(domain.Root, domain.Hope2Cmd, contract.Json2Hope, domain.Fbk2Json)
	natsResponder = nats.NewNatsResponder(initialize_nats.Bus, topics.NATS_HOPE_TOPIC, hopeHandler)
	// And
	natsResponder.Activate()

	// When
	f := initialize_nats.RequestByNats(hope)
	fbk := f.(*domain.Fbk)

	f = initialize_nats.RequestByNats(hope)
	//
	f = initialize_nats.RequestByNats(hope)
	// Then
	assert.NotNil(t, fbk)
	assert.Equal(t, model.Initialized, model.GreenhouseStatus(fbk.Status()))
}
