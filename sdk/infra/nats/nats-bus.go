package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/rgfaber/go-vesca/sdk/infra/nats/config"
	"log"
	"sync"
	"time"
)

var lock = &sync.Mutex{}

var singleNats *NatsBus

type INatsBus interface {
	Close()
	Publish(topic string, data []byte)
	Listen(topic string, facts chan []byte)
	Respond(topic string, hopes chan *nats.Msg)
	Request(topic string, data []byte, timeout time.Duration) []byte
}

func ImplementsINatsBus(b INatsBus) bool {
	return true
}

type NatsBus struct {
	conn *nats.Conn
}

func (b *NatsBus) Publish(topic string, data []byte) {
	err := b.conn.Publish(topic, data)
	if err != nil {
		log.Fatal(err)
	}
	err = b.conn.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

func (b *NatsBus) Listen(topic string, facts chan []byte) {
	go func(f chan []byte) {
		sub, err := b.conn.Subscribe(topic,
			func(msg *nats.Msg) {
				f <- msg.Data
			})
		if err != nil {
			log.Fatal(err)
		}
		err = b.conn.Flush()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Listening for NATS Fact [%s]\n", sub.Subject)
		select {}
	}(facts)
}

func (b *NatsBus) Respond(topic string, hopes chan *nats.Msg) {
	go func(c chan *nats.Msg) {
		sub, e := b.conn.Subscribe(topic, func(msg *nats.Msg) {
			c <- msg
		})
		if e != nil {
			log.Fatal(e)
		}
		err := b.conn.Flush()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Responding to NATS Hope [%s]\n", sub.Subject)
	}(hopes)
}

func (b *NatsBus) Request(topic string, data []byte, timeout time.Duration) []byte {
	msg, err := b.conn.Request(topic, data, timeout)
	if err != nil {
		log.Fatal(err)
	}
	return msg.Data
}

func (b *NatsBus) Connection() *nats.Conn {
	return b.conn
}

func SingletonNatsBus(cfg config.INatsConfig) INatsBus {
	if singleNats == nil {
		lock.Lock()
		defer lock.Unlock()
		conn, err := nats.Connect(
			cfg.NATSUrl(),
			nats.UserInfo(cfg.NATSUser(), cfg.NATSPwd()))
		if err != nil {
			log.Fatal(err)
			defer conn.Close()
		}
		if conn == nil {
			panic(err)
		}
		singleNats = &NatsBus{
			conn: conn,
		}
	}
	return singleNats
}

func TransientNatsBus(cfg config.INatsConfig) INatsBus {
	conn, err := nats.Connect(
		cfg.NATSUrl(),
		nats.UserInfo(cfg.NATSUser(), cfg.NATSPwd()))
	if err != nil {
		log.Fatal(err)
		defer conn.Close()
	}
	if conn == nil {
		panic(err)
	}
	return &NatsBus{
		conn: conn,
	}
}

func (b *NatsBus) Close() {
	b.conn.Close()
}
