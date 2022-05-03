package infra

import (
	"github.com/nats-io/nats.go"
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"log"
)

type INatsBus interface {
	Close()
	Publish(topic string, data []byte)
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

func (b *NatsBus) Connection() *nats.Conn {
	return b.conn
}

func NewNatsBus(cfg *config.Config) *NatsBus {
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
