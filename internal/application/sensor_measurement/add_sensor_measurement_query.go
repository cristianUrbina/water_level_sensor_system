package sensormeasurementapp

import (
	"time"

	"github.com/google/uuid"
)

type AddSensorMeasurementQuery struct {
	SensorID   uuid.UUID
	MeasuredAt time.Time
	Value      float64
	Type       string
}
