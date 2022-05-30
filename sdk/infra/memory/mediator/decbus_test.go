package mediator

import (
	"fmt"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/mediator/bogus"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	b := SingletonDECBus()
	assert.NotNil(t, b)
}

func TestSubscribe(t *testing.T) {
	b := SingletonDECBus()
	assert.NotNil(t, b)
	result := ""
	err := b.Subscribe(bogus.TestTopic, func(msg *bogus.TestMsg) {
		result = fmt.Sprintf("Received Message [%+v] on Topic [%+v]", msg.Content, bogus.TestTopic)
	})
	if err != nil {
		log.Fatal(err)
	}
	b.Publish(bogus.TestTopic, bogus.NewTestMsg("Hello There!"))
	assert.True(t, strings.Contains(result, "Hello There!"))
}

func TestSubscribeOnce(t *testing.T) {
	b := SingletonDECBus()
	result := ""
	assert.NotNil(t, b)
	err := b.SubscribeOnce(bogus.TestTopic, func(msg *bogus.TestMsg) {
		result = fmt.Sprintf("Received Message [%+v] on Topic [%+v]", msg.Content, bogus.TestTopic)
	})
	if err != nil {
		log.Fatal(err)
	}
	b.Publish(bogus.TestTopic, bogus.NewTestMsg("First Message"))
	b.Publish(bogus.TestTopic, bogus.NewTestMsg("Second Message"))
	assert.False(t, strings.Contains(result, "Second Message"))
}

func TestHasCallback(t *testing.T) {
	b := TransientDECBus()
	assert.NotNil(t, b)
	result := b.HasCallback(bogus.TestTopic)
	assert.False(t, result)
	err := b.SubscribeOnce(bogus.TestTopic, func(msg *bogus.TestMsg) {})
	assert.Nil(t, err)
	result = b.HasCallback(bogus.TestTopic)
	assert.True(t, result)
}

func helloThere() string {
	return "Hello There"
}

func TestUnsubscribe(t *testing.T) {
	b := TransientDECBus()
	assert.NotNil(t, b)
	err := b.SubscribeOnce(bogus.TestTopic, helloThere)
	assert.Nil(t, err)
	result := b.HasCallback(bogus.TestTopic)
	assert.True(t, result)
	err = b.Unsubscribe(bogus.TestTopic, helloThere)
	assert.Nil(t, err)
	result = b.HasCallback(bogus.TestTopic)
	assert.False(t, result)
}

func TestPublish(t *testing.T) {
	b := SingletonDECBus()
	assert.NotNil(t, b)
	result := ""
	err := b.SubscribeOnce(bogus.TestTopic, func(msg *bogus.TestMsg) {
		result = fmt.Sprintf(msg.Content)
	})
	assert.Nil(t, err)
	b.Publish(bogus.TestTopic, bogus.NewTestMsg("hello"))
	assert.Equal(t, result, "hello")

}

func TestSubscribeAsync(t *testing.T) {
	b := SingletonDECBus()
	assert.NotNil(t, b)
	result := ""
	err := b.SubscribeAsync(bogus.TestTopic, func(msg *bogus.TestMsg) {
		time.Sleep(6 * time.Second)
		result = msg.Content
	}, false)
	assert.Nil(t, err)
	b.Publish(bogus.TestTopic, bogus.NewTestMsg("hello"))
	fmt.Println("waiting for callback")
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("Waiting %+v sec\n", i+1)
	}
	b.WaitAsync()
	assert.Equal(t, "hello", result)
}

func longHelloThere(msg *bogus.TestMsg) {
	fmt.Println("Received msg:", msg.Content)
	time.Sleep(10 * time.Second)
}

func TestSubscribeOnceAsync(t *testing.T) {
	b := TransientDECBus()
	assert.NotNil(t, b)
	err := b.SubscribeOnceAsync(bogus.TestTopic, longHelloThere)
	assert.Nil(t, err)
	assert.True(t, b.HasCallback(bogus.TestTopic))
	for i := 0; i < 11; i++ {
		time.Sleep(1 * time.Second)
		if i == 3 {
			b.Publish(bogus.TestTopic, bogus.NewTestMsg("hello"))
			b.WaitAsync()
		}
		fmt.Printf("On second %+v => HasCallback is %+v\n", i, b.HasCallback(bogus.TestTopic))
	}
	assert.False(t, b.HasCallback(bogus.TestTopic))
}

func TestWaitAsync(t *testing.T) {
	b := SingletonDECBus()
	assert.NotNil(t, b)
	err := b.SubscribeAsync(bogus.TestTopic, longHelloThere, true)
	assert.Nil(t, err)
	b.Publish(bogus.TestTopic, bogus.NewTestMsg("Testing WaitAsync"))
	b.WaitAsync()
}
