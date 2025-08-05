package centrifugox

import (
	"context"
	"encoding/json"

	"github.com/centrifugal/centrifuge-go"
	"github.com/hoomanfr/harry/golib/config"
)

type Message struct {
	Data []byte `json:"data"`
}

type Broker struct {
	client *centrifuge.Client
}

func NewBroker(cfg *config.Config) (*Broker, error) {
	client := centrifuge.NewJsonClient(cfg.CentrifugoUrl, centrifuge.Config{
		Token: cfg.CentrifugoToken,
	})

	return &Broker{
		client: client,
	}, nil
}

func (b *Broker) Connect() error {
	return b.client.Connect()
}

func (b *Broker) Disconnect() error {
	b.client.Close()
	return nil
}

func (b *Broker) Publish(ctx context.Context, channel string, data any) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	msg := Message{
		Data: dataBytes,
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	_, err = b.client.Publish(ctx, channel, msgBytes)
	if err != nil {
		return err
	}

	return nil
}
