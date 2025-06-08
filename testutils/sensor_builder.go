package testutils

import (
	"cristianUrbina/water_level_sensor_system/internal/domain/sensordm"
	"cristianUrbina/water_level_sensor_system/internal/domain/tank"

	"github.com/google/uuid"
)

func NewSensorBuilder() *SensorBuilder {
	return &SensorBuilder{}
}

type SensorBuilder struct {
	id uuid.UUID
	name        string
	description string
	tank        *tank.Tank
}

func (s *SensorBuilder) WithID(id uuid.UUID) *SensorBuilder {
	s.id = id
	return s
}

func (s *SensorBuilder) WithDefaultValues() *SensorBuilder {
	s.name = "testing_sensor"
	s.description = "sensor used for testing"
	s.tank = tank.NewTank()
	return s
}

func (s *SensorBuilder) Build() (*sensordm.Sensor, error) {
	return sensordm.GetWaterLevelSensor(
		s.id,
		s.name,
		s.description,
		s.tank,
	)
}
