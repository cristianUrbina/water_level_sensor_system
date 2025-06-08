package common

import (
	"fmt"

	"github.com/google/uuid"
)

func NewNotFoundSensorError(sensorID uuid.UUID) *NotFoundSensor {
	return &NotFoundSensor{
		Msg: fmt.Sprintf("Not found sensor with id: %s ", sensorID),
	}
}

type NotFoundSensor struct {
	Msg string
}

func (e *NotFoundSensor) Error() string {
	return e.Msg
}
