package sensormeasurement

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

func NewSensorMeasurement(id, sensorID uuid.UUID, measuredAt time.Time, value float64, measType string) (*SensorMeasurement, error) {
	if id == uuid.Nil {
		return nil, errors.New("cannot have nil id")
	}
	return &SensorMeasurement{
		ID: id,
		SensorID: sensorID,
		MeasuredAt: measuredAt,
		Value: value,
		Type: measType,
	}, nil
}

type SensorMeasurement struct {
	ID         uuid.UUID `gorm:"type:varchar(36);primaryKey"`
	SensorID   uuid.UUID `gorm:"type:varchar(36)"`
	MeasuredAt time.Time
	Value      float64
	Type       string
}
