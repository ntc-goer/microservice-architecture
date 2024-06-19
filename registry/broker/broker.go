package broker

import (
	"github.com/nats-io/nats.go"
)

type Broker struct {
	Client *nats.Conn
}

func NewBroker() *Broker {
	return &Broker{}
}

func (q *Broker) Connect(addr string) error {
	nc, err := nats.Connect(addr)
	if err != nil {
		return err
	}
	q.Client = nc
	return nil
}

func (q *Broker) Close() {
	q.Client.Close()
}
func (q *Broker) Publish(subject string, data string) error {
	return q.Client.Publish(subject, []byte(data))
}

func (q *Broker) Subscribe(subject string, handler func(msg string)) error {
	_, err := q.Client.Subscribe(subject, func(msg *nats.Msg) {
		handler(string(msg.Data))
	})
	return err
}

func (q *Broker) QueueSubscribe(subject string, queueName string, handler func(msg string)) error {
	_, err := q.Client.QueueSubscribe(subject, queueName, func(msg *nats.Msg) {
		handler(string(msg.Data))
	})
	return err
}
