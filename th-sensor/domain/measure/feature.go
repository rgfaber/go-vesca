package measure

import (
	"fmt"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"log"
	"time"
)

type Feature struct {
	state *model.Root
	bus   dec.IDECBus
}

func (f *Feature) Listen() {
	err := f.bus.Subscribe(EVT_TOPIC, func(evt *Evt) {
		f.apply(evt)
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("measure.Feature -> Listening for [%+v]\n", EVT_TOPIC)
}

func (f *Feature) Respond() {
	err := f.bus.Subscribe(CMD_TOPIC, func(cmd *Cmd) {
		fmt.Printf("Received \n [%+v] \n on [%+v]\n", cmd.Measurement, CMD_TOPIC)
		f.exec(cmd)
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("measure.Feature -> Responding to [%+v]\n", CMD_TOPIC)
}

func NewFeature(state *model.Root, bus dec.IDECBus) *Feature {
	return &Feature{
		state: state,
		bus:   bus,
	}
}

func (f *Feature) Run() {
	go f.Listen()
	go f.Respond()
	fmt.Println("measure.Feature is Up!")
	go f.measure()
	//	select {}
}

func (f *Feature) apply(evt *Evt) {
	f.state.Measurement = evt.Measurement
	f.state.Status.Set(model.Measuring)
	fmt.Println("New Measurement:")
	fmt.Println(f.state.Measurement)
}

func (f *Feature) exec(cmd *Cmd) {
	if cmd == nil {
		log.Fatal(cmd)
		return
	}
	f.raise(NewEvt(cmd.Measurement))
}

func (f *Feature) raise(evt *Evt) {
	f.bus.Publish(EVT_TOPIC, evt)
}

func (f *Feature) measure() {
	for {
		fmt.Println("Triggering Measurement in 2s")
		time.Sleep(2 * time.Second)
		nm := model.NewMeasurement(f.state.Measurement.Temperature+.2, f.state.Measurement.Humidity+1)
		f.bus.Publish(CMD_TOPIC, NewCmd(nm))
	}
}
