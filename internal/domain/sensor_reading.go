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
		return SensorReading{}, ErrInvalidTimestamp
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
	ID         uuid.UUID `gorm:"type:varchar(36);primaryKey"`
	SensorID   uuid.UUID
	Sensor     Sensor `gorm:"foreignKey:SensorID;references:ID"`
	Capability Capability
	Value      float64
	Unit       string
	Timestamp  time.Time
}
