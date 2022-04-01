package infra

import (
	"github.com/nats-io/nats.go"
	config2 "github.com/rgfaber/go-vesca/th-sensor/config"
	"log"
)

type IBus interface {
	Close()
}

type Bus struct {
	cfg  *config2.Config
	conn *nats.Conn
}

func NewBus(cfg *config2.Config) *Bus {
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
	return &Bus{
		cfg:  cfg,
		conn: conn,
	}
}

func (b *Bus) Close() {
	b.conn.Close()
}
