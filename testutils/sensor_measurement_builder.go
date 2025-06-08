package testutils

import (
	"time"

	sensormeasurement "cristianUrbina/water_level_sensor_system/internal/domain/sensor_measurement"

	"github.com/google/uuid"
)

func NewSensorMeasurmentBuilder() *SensorMeasurmentBuilder {
	return &SensorMeasurmentBuilder{}
}

type SensorMeasurmentBuilder struct {
	id         uuid.UUID
	sensorID   uuid.UUID
	measuredAt time.Time
	value      float64
	mesType    string
}

func (s *SensorMeasurmentBuilder) WithSensorID(id uuid.UUID) *SensorMeasurmentBuilder {
	s.sensorID = id
	return s
}

func (s *SensorMeasurmentBuilder) WithDefaultValues() *SensorMeasurmentBuilder {
	s.id = uuid.New()
	s.sensorID = uuid.New()
	s.measuredAt = time.Now()
	s.value = float64(0.320)
	s.mesType = "liquid_level"
	return s
}

func (s *SensorMeasurmentBuilder) Build() (*sensormeasurement.SensorMeasurement, error) {
	return sensormeasurement.NewSensorMeasurement(
		s.id,
		s.sensorID,
		s.measuredAt,
		s.value,
		s.mesType)
}
