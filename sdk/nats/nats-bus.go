package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/sdk/nats/config"
	"log"
)

type INatsBus interface {
	Close()
	Publish(topic string, data []byte)
	Listen(topic string, handler dec.IFactHandler)
	Respond(topic string, handler dec.IHopeHandler)
	Connection() *nats.Conn
}

func ImplementsINatsBus(b INatsBus) bool {
	return true
}

type NatsBus struct {
	conn *nats.Conn
}

func (b *NatsBus) Publish(topic string, data []byte) {
	b.conn.Publish(topic, data)
}

func (b *NatsBus) Listen(topic string, handler dec.IFactHandler) {
	b.conn.Subscribe(topic, func(msg *nats.Msg) {
		handler.Handle(msg.Data)
	})
}

func (b *NatsBus) Respond(topic string, handler dec.IHopeHandler) {
	b.conn.Subscribe(topic, func(msg *nats.Msg) {
		fbk := handler.Handle(msg.Data)
		err := msg.Respond(fbk)
		if err != nil {
			log.Fatal(err)
		}
	})
}

func (b *NatsBus) Connection() *nats.Conn {
	return b.conn
}

func NewNatsBus(cfg config.INatsConfig) *NatsBus {
	conn, err := nats.Connect(
		cfg.NATSUrl(),
		nats.UserInfo(cfg.NATSUser(), cfg.NATSPwd()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
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
