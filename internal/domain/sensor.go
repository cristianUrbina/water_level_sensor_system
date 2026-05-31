package domain

import (
	"time"

	"github.com/google/uuid"
)

type Capability string

type MeasurementProfile string

const (
	ProfileTankLevel MeasurementProfile = "tank_level"
	ProfileDistance  MeasurementProfile = "distance"
	ProfileWeather   MeasurementProfile = "weather"
)

type Sensor struct {
	ID uuid.UUID

	Metadata SensorMetadata

	Capabilities []Capability

	Profile MeasurementProfile

	CreatedAt time.Time
	UpdatedAt time.Time
}

type SensorMetadata struct {
	Name string

	Description string

	HardwareModel string

	CalibrationOffset float64
}
