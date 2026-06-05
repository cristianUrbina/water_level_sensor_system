package mqtt

import (
	"fmt"
	"log"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
)

type PahoClient struct {
	client paho.Client
}

type Config struct {
	BrokerURL string
	ClientID  string
	Username  string
	Password  string
}

func NewPahoClient(cfg Config) (*PahoClient, error) {
	log.Println("Connecting to MQTT broker", cfg)
	opts := paho.NewClientOptions()

	opts.AddBroker(cfg.BrokerURL)
	opts.SetClientID(cfg.ClientID)

	if cfg.Username != "" {
		opts.SetUsername(cfg.Username)
		opts.SetPassword(cfg.Password)
	}

	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)

	opts.OnConnect = func(c paho.Client) {
		log.Println("MQTT connected")
	}

	opts.OnConnectionLost = func(c paho.Client, err error) {
		log.Printf("MQTT connection lost: %v", err)
	}

	client := paho.NewClient(opts)

	token := client.Connect()

	// ✅ Avoid infinite wait
	if !token.WaitTimeout(5 * time.Second) {
		return nil, fmt.Errorf("mqtt connection timeout")
	}

	// ✅ Check token error
	if err := token.Error(); err != nil {
		return nil, fmt.Errorf("connect mqtt broker: %w", err)
	}

	// ✅ Final safety check
	if !client.IsConnected() {
		return nil, fmt.Errorf("mqtt client not connected")
	}

	log.Println("MQTT connection established successfully ✅")

	return &PahoClient{
		client: client,
	}, nil
}

func (c *PahoClient) Subscribe(
	topic string,
	qos byte,
	handler MessageHandler,
) error {
	token := c.client.Subscribe(
		topic,
		qos,
		func(_ paho.Client, msg paho.Message) {
			_ = handler.Handle(Message{
				Topic:   topic,
				Payload: msg.Payload(),
			})
		},
	)
	token.Wait()
	return token.Error()
}
