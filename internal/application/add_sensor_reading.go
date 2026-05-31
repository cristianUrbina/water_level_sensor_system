package application

import (
	"context"
	"errors"
	"time"

	"cristianUrbina/water_level_sensor_system/internal/domain"

	"github.com/google/uuid"
)

type IAddSensorReadingHandler interface {
	Handle(ctx context.Context, query AddSensorReadingQuery) error
}

type AddSensorReadingQuery struct {
	SensorID   uuid.UUID
	Capability string
	Value      float64
	Unit       string
	Timestamp  time.Time
}

type AddSensorReadingHandler struct {
	Sensors  domain.ISensorRepository
	Readings domain.IReadingRepository
}

func NewAddSensorReadingHandler(readings domain.IReadingRepository, sensors domain.ISensorRepository) AddSensorReadingHandler {
	return AddSensorReadingHandler{
		Sensors:  sensors,
		Readings: readings,
	}
}

func (a AddSensorReadingHandler) Handle(ctx context.Context, query AddSensorReadingQuery) error {
	sensor, err := a.Sensors.GetByID(ctx, query.SensorID)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return ErrSensorNotFound
		}
		return err
	}
	capability, err := domain.ParseCapability(query.Capability)
	if err != nil {
		return err
	}
	reading, err := domain.NewSensorReading(sensor.ID, capability, query.Value, query.Unit, query.Timestamp)
	if err != nil {
		return ErrInvalidEntity
	}
	return a.Readings.Add(ctx, reading)
}
