package mqtt

import (
	"context"
	"encoding/json"
	"time"

	"cristianUrbina/water_level_sensor_system/internal/application"
	pkg "cristianUrbina/water_level_sensor_system/pkg/errors"

	"github.com/google/uuid"
	"github.com/mehdihadeli/go-mediatr"
)

type ReadingHandler struct{}

func NewReadingHandler() *ReadingHandler {
	return &ReadingHandler{}
}

func (r *ReadingHandler) Handle(msg Message) error {
	var dto sensorReadingDTO
	if err := json.Unmarshal(msg.Payload, &dto); err != nil {
		return err
	}
	sensorID, err := uuid.Parse(dto.SensorID)
	if err != nil {
		return pkg.ErrInvalidUUID
	}
	timestamp, err := time.Parse(time.RFC3339, dto.Timestamp)
	if err != nil {
		return pkg.ErrInvalidTimestamp
	}
	q := application.AddSensorReadingQuery{
		SensorID:   sensorID,
		Capability: dto.Capability,
		Value:      dto.Value,
		Unit:       dto.Unit,
		Timestamp:  timestamp,
	}
	_, err = mediatr.Send[application.AddSensorReadingQuery, mediatr.Unit](context.Background(), q)
	return err
}
