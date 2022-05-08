package sdk

import (
	"fmt"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	b := NewDECBus()
	assert.NotNil(t, b)
}

func TestSubscribe(t *testing.T) {
	b := NewDECBus()
	assert.NotNil(t, b)
	result := ""
	err := b.Subscribe(dec.TestTopic, func(msg *dec.TestMsg) {
		result = fmt.Sprintf("Received Message [%+v] on Topic [%+v]", msg.Content, dec.TestTopic)
	})
	if err != nil {
		log.Fatal(err)
	}
	b.bus.Publish(dec.TestTopic, dec.NewTestMsg("Hello There!"))
	assert.True(t, strings.Contains(result, "Hello There!"))
}

func TestSubscribeOnce(t *testing.T) {
	b := NewDECBus()
	result := ""
	assert.NotNil(t, b)
	err := b.SubscribeOnce(dec.TestTopic, func(msg *dec.TestMsg) {
		result = fmt.Sprintf("Received Message [%+v] on Topic [%+v]", msg.Content, dec.TestTopic)
	})
	if err != nil {
		log.Fatal(err)
	}
	b.bus.Publish(dec.TestTopic, dec.NewTestMsg("First Message"))
	b.bus.Publish(dec.TestTopic, dec.NewTestMsg("Second Message"))
	assert.False(t, strings.Contains(result, "Second Message"))
}

func TestHasCallback(t *testing.T) {
	b := NewDECBus()
	assert.NotNil(t, b)
	result := b.HasCallback(dec.TestTopic)
	assert.False(t, result)
	err := b.SubscribeOnce(dec.TestTopic, func(msg *dec.TestMsg) {})
	assert.Nil(t, err)
	result = b.HasCallback(dec.TestTopic)
	assert.True(t, result)
}

func helloThere() string {
	return "Hello There"
}

func TestUnsubscribe(t *testing.T) {
	b := NewDECBus()
	assert.NotNil(t, b)
	err := b.bus.SubscribeOnce(dec.TestTopic, helloThere)
	assert.Nil(t, err)
	result := b.bus.HasCallback(dec.TestTopic)
	assert.True(t, result)
	err = b.Unsubscribe(dec.TestTopic, helloThere)
	assert.Nil(t, err)
	result = b.bus.HasCallback(dec.TestTopic)
	assert.False(t, result)
}

func TestPublish(t *testing.T) {
	b := NewDECBus()
	assert.NotNil(t, b)
	result := ""
	err := b.bus.SubscribeOnce(dec.TestTopic, func(msg *dec.TestMsg) {
		result = fmt.Sprintf(msg.Content)
	})
	assert.Nil(t, err)
	b.Publish(dec.TestTopic, dec.NewTestMsg("hello"))
	assert.Equal(t, result, "hello")

}

func TestSubscribeAsync(t *testing.T) {
	b := NewDECBus()
	assert.NotNil(t, b)
	result := ""
	err := b.SubscribeAsync(dec.TestTopic, func(msg *dec.TestMsg) {
		time.Sleep(6 * time.Second)
		result = msg.Content
	}, false)
	assert.Nil(t, err)
	b.bus.Publish(dec.TestTopic, dec.NewTestMsg("hello"))
	fmt.Println("waiting for callback")
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("Waiting %+v sec\n", i+1)
	}
	b.bus.WaitAsync()
	assert.Equal(t, "hello", result)
}

func longHelloThere(msg *dec.TestMsg) {
	fmt.Println("Received msg:", msg.Content)
	time.Sleep(10 * time.Second)
}

func TestSubscribeOnceAsync(t *testing.T) {
	b := NewDECBus()
	assert.NotNil(t, b)
	err := b.SubscribeOnceAsync(dec.TestTopic, longHelloThere)
	assert.Nil(t, err)
	assert.True(t, b.bus.HasCallback(dec.TestTopic))
	for i := 0; i < 11; i++ {
		time.Sleep(1 * time.Second)
		if i == 3 {
			b.bus.Publish(dec.TestTopic, dec.NewTestMsg("hello"))
			b.bus.WaitAsync()
		}
		fmt.Printf("On second %+v => HasCallback is %+v\n", i, b.bus.HasCallback(dec.TestTopic))
	}
	assert.False(t, b.bus.HasCallback(dec.TestTopic))
}

func TestWaitAsync(t *testing.T) {
	b := NewDECBus()
	assert.NotNil(t, b)
	err := b.bus.SubscribeAsync(dec.TestTopic, longHelloThere, true)
	assert.Nil(t, err)
	b.bus.Publish(dec.TestTopic, dec.NewTestMsg("Testing WaitAsync"))
	b.WaitAsync()
}
