package domain_test

import (
	"testing"
	"time"

	"cristianUrbina/water_level_sensor_system/internal/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSensorReadingCreation(t *testing.T) {
	sensorID := uuid.New()
	value := 23.3
	unit := "cm"
	timestamp := time.Now()
	sensorReading, err := domain.NewSensorReading(sensorID, domain.CapabilityDistance, value, unit, timestamp)

	assert.NoError(t, err)
	assert.Equal(t, sensorReading.SensorID, sensorID)
	assert.Equal(t, sensorReading.Capability, domain.CapabilityDistance)
	assert.Equal(t, sensorReading.Value, value)
	assert.Equal(t, sensorReading.Unit, unit)
	assert.Equal(t, sensorReading.Timestamp, timestamp)
}

func TestSensorReadingRequiresSensorID(t *testing.T) {
	_, err := domain.NewSensorReading(uuid.Nil, domain.CapabilityDistance, 23.3, "cm", time.Now())
	assert.Error(t, err)
}

func TestSensorReadingRequiresTimestamp(t *testing.T) {
	_, err := domain.NewSensorReading(uuid.New(), domain.CapabilityDistance, 23.3, "cm", time.Time{})
	assert.ErrorIs(t, err, domain.InvalidTimestampErr)
}
