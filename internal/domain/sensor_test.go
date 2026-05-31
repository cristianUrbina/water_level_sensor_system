package domain_test

import (
	"cristianUrbina/water_level_sensor_system/internal/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSensorCreation(t *testing.T) {
	sensor := domain.Sensor{
		ID: uuid.New(),

		Profile: domain.ProfileTankLevel,

		Capabilities: []domain.Capability{
			domain.CapabilityDistance,
		},

		Metadata: domain.SensorMetadata{
			Name:            "Tank Sensor 1",
			Description:     "Ultrasonic water tank sensor",
			HardwareModel:   "JSN-SR04T",
		},
	}

	assert.NotEqual(t, uuid.Nil, sensor.ID)

	assert.Equal(
		t,
		domain.ProfileTankLevel,
		sensor.Profile,
	)

	assert.Len(t, sensor.Capabilities, 1)

	assert.Contains(
		t,
		sensor.Capabilities,
		domain.CapabilityDistance,
	)

	assert.Equal(
		t,
		"Tank Sensor 1",
		sensor.Metadata.Name,
	)

	assert.Equal(
		t,
		"JSN-SR04T",
		sensor.Metadata.HardwareModel,
	)
}

func TestSensorSupportsMultipleCapabilities(
	t *testing.T,
) {
	sensor := domain.Sensor{
		ID: uuid.New(),

		Profile: domain.ProfileWeather,

		Capabilities: []domain.Capability{
			domain.CapabilityTemperature,
			domain.CapabilityHumidity,
		},
	}

	assert.Len(t, sensor.Capabilities, 2)

	assert.Contains(
		t,
		sensor.Capabilities,
		domain.CapabilityTemperature,
	)

	assert.Contains(
		t,
		sensor.Capabilities,
		domain.CapabilityHumidity,
	)
}

func TestDifferentProfilesCanReuseCapability(
	t *testing.T,
) {

	tankSensor := domain.Sensor{
		ID: uuid.New(),

		Profile: domain.ProfileTankLevel,

		Capabilities: []domain.Capability{
			domain.CapabilityDistance,
		},
	}

	distanceSensor := domain.Sensor{
		ID: uuid.New(),

		Profile: domain.ProfileDistance,

		Capabilities: []domain.Capability{
			domain.CapabilityDistance,
		},
	}

	assert.Equal(
		t,
		domain.CapabilityDistance,
		tankSensor.Capabilities[0],
	)

	assert.Equal(
		t,
		domain.CapabilityDistance,
		distanceSensor.Capabilities[0],
	)

	assert.NotEqual(
		t,
		tankSensor.Profile,
		distanceSensor.Profile,
	)
}

func TestSensorMetadata(t *testing.T) {
	metadata := domain.SensorMetadata{
		Name:            "Outdoor Sensor",
		Description:     "Environmental sensor",
		HardwareModel:   "BME280",
	}

	assert.Equal(
		t,
		"Outdoor Sensor",
		metadata.Name,
	)

	assert.Equal(
		t,
		"BME280",
		metadata.HardwareModel,
	)
}

