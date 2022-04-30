package infra

import (
	"github.com/nats-io/nats.go"
	config2 "github.com/rgfaber/go-vesca/th-sensor/config"
	"log"
)

type INatsBus interface {
	Close()
}

type NatsBus struct {
	cfg  *config2.Config
	conn *nats.Conn
}

func NewNatsBus(cfg *config2.Config) *NatsBus {
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
		cfg:  cfg,
		conn: conn,
	}
}

func (b *NatsBus) Close() {
	b.conn.Close()
}
