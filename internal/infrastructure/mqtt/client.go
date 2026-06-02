package mqtt

type Message struct {
	Topic string
	Payload []byte
}

type MessageHandler interface {
	Handle(Message) error
}

type Client interface {
	Subscribe(
		topic string,
		qos byte,
		callback MessageHandler,
	)
}
