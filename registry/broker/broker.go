package broker

import (
	"context"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel/propagation"
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
func (q *Broker) Publish(ctx context.Context, subject string, data []byte) error {
	header := make(nats.Header)

	// ADD THIS FOR TRACE CONTEXT PROPAGATION
	propagator := propagation.TraceContext{}
	propagator.Inject(ctx, propagation.HeaderCarrier(header))
	// *******
	return q.Client.PublishMsg(&nats.Msg{
		Subject: subject,
		Header:  header,
		Data:    data,
	})
}

func (q *Broker) Subscribe(ctx context.Context, subject string, handler func(ctx context.Context, msg string)) error {
	_, err := q.Client.Subscribe(subject, func(msg *nats.Msg) {
		propagator := propagation.TraceContext{}
		ctx = propagator.Extract(ctx, propagation.HeaderCarrier(msg.Header))
		handler(ctx, string(msg.Data))
	})
	return err
}

func (q *Broker) QueueSubscribe(ctx context.Context, subject string, queueName string, handler func(ctx context.Context, msg string)) error {
	_, err := q.Client.QueueSubscribe(subject, queueName, func(msg *nats.Msg) {
		propagator := propagation.TraceContext{}
		ctx = propagator.Extract(ctx, propagation.HeaderCarrier(msg.Header))
		handler(ctx, string(msg.Data))
	})
	return err
}
