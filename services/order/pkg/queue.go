package pkg

import (
	"github.com/nats-io/nats.go"
	"github.com/ntc-goer/microservice-examples/orderservice/config"
)

type MsgQueue struct {
	Config *config.Config
	Client *nats.Conn
}

func NewMsgQueue(cfg *config.Config) *MsgQueue {
	return &MsgQueue{
		Config: cfg,
	}
}

func (q *MsgQueue) Connect() error {
	nc, err := nats.Connect(q.Config.QueueHost)
	if err != nil {
		return err
	}
	q.Client = nc
	return nil
}

func (q *MsgQueue) Publish(subject string, data string) error {
	return q.Client.Publish(subject, []byte(data))
}

func (q *MsgQueue) Subscribe(subject string, handler func(msg string)) {
	for {
		q.Client.Subscribe(subject, func(msg *nats.Msg) {
			handler(string(msg.Data))
		})
	}
}
