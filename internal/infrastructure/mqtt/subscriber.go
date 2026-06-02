package mqtt

import (
	paho "github.com/eclipse/paho.mqtt.golang"
)

type Subscriber struct {
	client paho.Client
}
