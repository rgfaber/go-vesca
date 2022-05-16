package dec

import (
	"fmt"
	"github.com/rgfaber/go-vesca/sdk/bogus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func processMsg(med Mediator) {
	for {
		msg := <-med
		switch msg.Topic() {
		case bogus.TestTopic:
			{
				tm, ok := msg.(*bogus.TestMsg)
				if ok {
					fmt.Println("Received msg.Test! Content:", tm.Content)
				}
			}
		case bogus.QuitTopic:
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
			med <- bogus.NewTestMsg(fmt.Sprint(i))
		}
		med <- bogus.NewQuitMsg()
	}()
	processMsg(med)
}
