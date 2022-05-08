package dec

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	TestTopic = "topic.test"
	QuitTopic = "topic.quit"
)

// TestMsg
type TestMsg struct {
	Content string
}

func NewTestMsg(content string) *TestMsg {
	return &TestMsg{
		Content: content,
	}
}
func (msg *TestMsg) Topic() string {
	return TestTopic
}

// QuitMsg
type QuitMsg struct{}

func NewQuitMsg() *QuitMsg {
	return &QuitMsg{}
}
func (q *QuitMsg) Topic() string {
	return QuitTopic
}

func processMsg(med Mediator) {
	for {
		msg := <-med
		switch msg.Topic() {
		case TestTopic:
			{
				tm, ok := msg.(*TestMsg)
				if ok {
					fmt.Println("Received msg.Test! Content:", tm.Content)
				}
			}
		case QuitTopic:
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
			med <- NewTestMsg(fmt.Sprint(i))
		}
		med <- NewQuitMsg()
	}()
	processMsg(med)
}
