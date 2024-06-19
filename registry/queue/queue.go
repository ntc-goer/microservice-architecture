package queue

import (
	"github.com/nats-io/nats.go"
)

type MsgQueue struct {
	Client *nats.Conn
}

func NewMsgQueue() *MsgQueue {
	return &MsgQueue{}
}

func (q *MsgQueue) Connect(addr string) error {
	nc, err := nats.Connect(addr)
	if err != nil {
		return err
	}
	q.Client = nc
	return nil
}

func (q *MsgQueue) Close() {
	q.Client.Close()
}
func (q *MsgQueue) Publish(subject string, data string) error {
	return q.Client.Publish(subject, []byte(data))
}

func (q *MsgQueue) Subscribe(subject string, handler func(msg string)) error {
	_, err := q.Client.Subscribe(subject, func(msg *nats.Msg) {
		handler(string(msg.Data))
	})
	return err
}

func (q *MsgQueue) QueueSubscribe(subject string, queueName string, handler func(msg string)) error {
	_, err := q.Client.QueueSubscribe(subject, queueName, func(msg *nats.Msg) {
		handler(string(msg.Data))
	})
	return err
}
