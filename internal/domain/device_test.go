package domain_test

import (
	"testing"

	"cristianUrbina/water_level_sensor_system/internal/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeviceCreation(t *testing.T) {
	device := domain.Device{
		ID:   uuid.New(),
		Name: "Tank device",
		Metadata: domain.DeviceMetadata{
			Location:        "Home tank",
			HardwareModel:   "esp32",
			FirmwareVersion: "1.0.2",
			SerialNumber:    "D098740187",
		},
		Sensors: []domain.Sensor{
			{
				ID: uuid.New(),
				Metadata: domain.SensorMetadata{
					Name:              "UltraSonic Sensor",
					Description:       "Waterproof ultrasonic sensor",
					HardwareModel:     "JSN-SR04T",
					CalibrationOffset: 0.03,
				},
			},
		},
	}
	assert.NotEqual(t, uuid.Nil, device.ID)

	assert.Equal(t, "Tank device", device.Name)

	assert.NotNil(t, device.Metadata)

	assert.Len(t, device.Sensors, 1)
}

func TestDeviceSupportsMultipleSensors(t *testing.T) {
	device := domain.Device{
		ID:   uuid.New(),
		Name: "Tank device",
		Metadata: domain.DeviceMetadata{
			Location:        "Home tank",
			HardwareModel:   "esp32",
			FirmwareVersion: "1.0.2",
			SerialNumber:    "D098740187",
		},
		Sensors: []domain.Sensor{
			{
				ID: uuid.New(),
				Metadata: domain.SensorMetadata{
					Name:              "UltraSonic Sensor",
					Description:       "Waterproof ultrasonic sensor",
					HardwareModel:     "JSN-SR04T",
					CalibrationOffset: 0.03,
				},
			},
			{
				ID: uuid.New(),
				Metadata: domain.SensorMetadata{
					Name:              "Pressure Sensor 1",
					Description:       "Pressure Sensor",
					HardwareModel:     "Mps20n0040d-s",
					CalibrationOffset: 0.03,
				},
			},
		},
	}
	assert.Len(t, device.Sensors, 2)
	assert.Equal(t, device.Sensors[0].Metadata.Description, "Waterproof ultrasonic sensor")
	assert.Equal(t, device.Sensors[1].Metadata.Description, "Pressure Sensor")
}
