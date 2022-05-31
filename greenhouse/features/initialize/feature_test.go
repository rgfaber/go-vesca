package initialize

import (
	"fmt"
	"github.com/rgfaber/go-vesca/greenhouse/config/envars"
	initialize_contract "github.com/rgfaber/go-vesca/greenhouse/features/initialize/contract"
	initialize_domain "github.com/rgfaber/go-vesca/greenhouse/features/initialize/domain"
	"github.com/rgfaber/go-vesca/greenhouse/features/initialize/infra/memory"
	initialize_nats "github.com/rgfaber/go-vesca/greenhouse/features/initialize/infra/nats"
	"github.com/rgfaber/go-vesca/greenhouse/features/initialize/topics"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/greenhouse/model/bogus"
	core_mocks "github.com/rgfaber/go-vesca/sdk/core/mocks"
	"github.com/rgfaber/go-vesca/sdk/domain"
	"github.com/rgfaber/go-vesca/sdk/features"
	"github.com/rgfaber/go-vesca/sdk/infra"
	"github.com/rgfaber/go-vesca/sdk/infra/nats"
	"github.com/stretchr/testify/assert"
	"os"
	"sync"
	"testing"
)

func setEnv() {
	err := os.Setenv(envars.GO_VESCA_TH_SENSOR_ID, core_mocks.TEST_UUID)
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
	cmd := initialize_domain.NewCmd(bogus.BogusGreenhouseID,
		core_mocks.TEST_TRACE_ID,
		bogus.BogusDetails,
		bogus.BogusSettings,
		bogus.BogusSensor,
		bogus.BogusFan,
		bogus.BogusSprinkler)
	aggregate := initialize_domain.Root
	// When
	memory.Mediator.Subscribe(topics.EVT_TOPIC, func(evt domain.IEvt) {
		e := evt.(*initialize_domain.Evt)
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
	imp := features.ImplementsIFeature(f)
	// Then
	assert.True(t, imp)
}

func Test_Fact2Json(t *testing.T) {
	// Given
	fact := initialize_contract.NewFact(core_mocks.TEST_UUID, core_mocks.TEST_TRACE_ID, bogus.BogusGreenhouse)
	// When
	json := initialize_contract.Fact2Json(fact)
	// Then
	assert.NotEmpty(t, json)
}

func Test_Hope2Json(t *testing.T) {
	// Given
	hope := initialize_contract.NewHope(bogus.BogusGreenhouseID.Id(),
		core_mocks.TEST_TRACE_ID,
		bogus.BogusDetails,
		bogus.BogusSettings,
		bogus.BogusSensor,
		bogus.BogusFan,
		bogus.BogusSprinkler)
	// When
	json := initialize_contract.Hope2Json(hope)
	// Then
	assert.NotEmpty(t, json)
}

func Test_Hope2Cmd(t *testing.T) {
	// Given
	hope := initialize_contract.NewHope(bogus.BogusGreenhouseID.Id(),
		core_mocks.TEST_TRACE_ID,
		bogus.BogusDetails,
		bogus.BogusSettings,
		bogus.BogusSensor,
		bogus.BogusFan,
		bogus.BogusSprinkler)
	// When
	cmd := initialize_contract.Hope2Cmd(hope)
	c := cmd.(*initialize_domain.Cmd)
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
	evt := initialize_domain.NewEvt(bogus.BogusGreenhouseID, core_mocks.TEST_TRACE_ID, bogus.BogusGreenhouse)
	testHandler := func(fact domain.IFact) {
		lock.Lock()
		defer lock.Unlock()
		factHandled = true
		f := fact.(*initialize_contract.Fact)
		fmt.Printf("Handled Fact #%+v => [%+v]\n", i, *f)
		i++
	}
	factHandler := infra.NewFactHandler(testHandler, initialize_contract.Json2Fact)

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
	id := bogus.BogusGreenhouseID.Id()
	hope := initialize_contract.NewHope(id, core_mocks.TEST_TRACE_ID, bogus.BogusDetails, bogus.BogusSettings, bogus.BogusSensor, bogus.BogusFan, bogus.BogusSprinkler)
	hopeHandler := infra.NewHopeHandler(initialize_domain.Root, initialize_contract.Hope2Cmd, initialize_contract.Json2Hope, initialize_contract.Fbk2Json)
	natsResponder := nats.NewNatsResponder(initialize_nats.Bus, topics.NATS_HOPE_TOPIC, hopeHandler)
	// And
	natsResponder.Activate()

	// When
	f := initialize_nats.RequestByNats(hope)
	fbk := f.(*initialize_domain.Fbk)

	f = initialize_nats.RequestByNats(hope)
	//
	f = initialize_nats.RequestByNats(hope)
	// Then
	assert.NotNil(t, fbk)
	assert.Equal(t, model.Initialized, model.GreenhouseStatus(fbk.Status()))
}
