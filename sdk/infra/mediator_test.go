package infra

import (
	"fmt"
	bogus2 "github.com/rgfaber/go-vesca/sdk/infra/memory/mediator/bogus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func processMsg(med Mediator) {
	for {
		msg := <-med
		switch msg.Topic() {
		case bogus2.TestTopic:
			{
				tm, ok := msg.(*bogus2.TestMsg)
				if ok {
					fmt.Println("Received msg.Test! Content:", tm.Content)
				}
			}
		case bogus2.QuitTopic:
			{
				fmt.Println("Received msg.Quit")
				return
			}
		}
	}
}

func TestNewMediator(t *testing.T) {
	m := NewMediator(10)
	assert.NotNil(t, m)
}

func TestIfWeCanPutDifferentMessagesOnTheMediator(t *testing.T) {
	med := NewMediator(10)
	go func() {
		for i := 0; i < 10; i++ {
			med <- bogus2.NewTestMsg(fmt.Sprint(i))
		}
		med <- bogus2.NewQuitMsg()
	}()
	processMsg(med)
}
