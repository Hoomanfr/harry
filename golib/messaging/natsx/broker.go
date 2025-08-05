package natsx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hoomanfr/harry/golib/config"
	"github.com/nats-io/nats.go"
)

type Message struct {
	Data []byte `json:"data"`
}

type Broker struct {
	urls          string
	connection    *nats.Conn
	domain        string
	service       string
	maxReconnect  int
	reconnectWait int
}

func NewBroker(cfg config.Config, domain string, service string) (*Broker, error) {
	if domain == "" {
		return nil, errors.New("domain is empty")
	}
	if service == "" {
		return nil, errors.New("service is empty")
	}
	return &Broker{
		urls:    cfg.NatsUrls,
		domain:  domain,
		service: service,
	}, nil
}

func (b *Broker) Connect() error {
	nc, err := nats.Connect(b.urls, nats.MaxReconnects(b.maxReconnect), nats.ReconnectWait(time.Duration(b.reconnectWait)*time.Second))
	if err != nil {
		return err
	}
	b.connection = nc
	return nil
}

func (b *Broker) Disconnect() error {
	return b.connection.Drain()
}

func (b *Broker) Publish(topic string, data any) error {
	if topic == "" {
		return errors.New("publish topic is empty")
	}
	if data == nil {
		return errors.New("publish data is nil")
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	msg := &Message{
		Data: dataBytes,
	}
	msgJson, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return b.connection.Publish(fmt.Sprintf("%s.%s.%s", b.domain, b.service, topic), msgJson)
}

type Subscriber struct {
	subscriberName string
	broker         *Broker
}

func NewSubscriber(broker *Broker) *Subscriber {
	return &Subscriber{
		subscriberName: fmt.Sprintf("%s-%s", broker.domain, broker.service),
		broker:         broker,
	}
}

func (s *Subscriber) Subscribe(ctx context.Context, domain string, service string, topic string, handler func(ctx context.Context, msg Message) error) error {
	msgs := make(chan *nats.Msg)
	subject := fmt.Sprintf("%s.%s.%s", domain, service, topic)
	queueName := fmt.Sprintf("%s-%s", s.subscriberName, strings.ReplaceAll(subject, ".", "-"))
	sub, err := s.broker.connection.QueueSubscribeSyncWithChan(subject, queueName, msgs)
	if err != nil {
		return err
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				err := sub.Unsubscribe()
				if err != nil {
					fmt.Printf("failed to unsubscribe from subject %s: %v\r\n", subject, err)
				}
				return
			case msg := <-msgs:
				var data Message
				err := json.Unmarshal(msg.Data, &data)
				if err != nil {
					fmt.Printf("failed to unmarshal message for subject %s: %v\r\n", msg.Subject, err)
					continue
				}
				err = handler(ctx, data)
				if err != nil {
					fmt.Printf("failed to handle message for subject %s: %v\r\n", msg.Subject, err)
				}
			}
		}
	}()
	return nil
}
