package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

func NewSensorReading(sensorID uuid.UUID, capability Capability, value float64, unit string, timestamp time.Time) (SensorReading, error) {
	if sensorID == uuid.Nil {
		return SensorReading{}, errors.New("sensor id is required")
	}

	if timestamp.IsZero() {
		return SensorReading{}, InvalidTimestampErr
	}

	return SensorReading{
		SensorID:   sensorID,
		Capability: capability,
		Value:      value,
		Unit:       unit,
		Timestamp:  timestamp,
	}, nil
}

type SensorReading struct {
	SensorID   uuid.UUID
	Capability Capability
	Value      float64
	Unit       string
	Timestamp  time.Time
}
