package ports

type SensorReadingEventConsumer interface {
	Start(handler func(topic string, payload []byte)) error
}
