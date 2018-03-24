package pubsub

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
)

const envRabbitmqDialUrl = "RABBITMQ_DIAL_STRING"

type PubSub struct {
	conn                 *amqp.Connection
	channel              *amqp.Channel
	initialized          bool
	initializedExchanges map[string]bool
}

var Std = newPubSub()

func newPubSub() PubSub {
	return PubSub{
		initializedExchanges: map[string]bool{},
	}
}

func (p *PubSub) connect() error {
	var err error
	p.conn, err = amqp.Dial(os.Getenv(envRabbitmqDialUrl))
	if err != nil {
		return fmt.Errorf("error dialing amqp: %s", err)
	}

	p.channel, err = p.conn.Channel()
	if err != nil {
		p.conn.Close()
		return fmt.Errorf("error creating channel: %s", err)
	}

	return nil
}

func (p *PubSub) initExchange(exchange string) error {
	return p.channel.ExchangeDeclare(exchange, "fanout", true, false, false, false, nil)
}

func (p *PubSub) Pub(exchange string, msg []byte) error {
	if !p.initialized {
		if err := p.connect(); err != nil {
			return fmt.Errorf("failed to init rabbitmq: %s", err)
		}
		p.initialized = true
	}

	if !p.initializedExchanges[exchange] {
		err := p.initExchange(exchange)
		if err != nil {
			return fmt.Errorf("error initializing exchange: %s", err)
		}
		p.initializedExchanges[exchange] = true
	}

	err := p.channel.Publish(exchange, "", false, false, amqp.Publishing{Body: msg})
	if err != nil {
		return fmt.Errorf("error sending message: %s", err)
	}

	return nil
}
