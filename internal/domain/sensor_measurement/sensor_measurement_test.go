package sensormeasurement

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSensorMeasurementConstructorShouldReturnSensorMeasurement(t *testing.T) {
	id := uuid.New()
	sensorID := uuid.New()
	measuredAt := time.Now()
	value := float64(0.320)
	mesType := "liquid_level"

	mes, err := NewSensorMeasurement(id, sensorID, measuredAt, value, mesType)

	assert.NotNil(t, mes)
	assert.Nil(t, err)
	assert.Equal(t, id, mes.ID)
	assert.Equal(t, sensorID, mes.SensorID)
	assert.Equal(t, measuredAt, mes.MeasuredAt)
	assert.Equal(t, value, mes.Value)
	assert.Equal(t, mesType, mes.Type)
}

func TestSensorMeasurementConstructorShouldReturnErrorWhenPassingInvalidID(t *testing.T) {
	id := uuid.Nil
	sensorID := uuid.New()
	measuredAt := time.Now()
	value := float64(0.320)
	mesType := "liquid_level"

	mes, err := NewSensorMeasurement(id, sensorID, measuredAt, value, mesType)

	assert.Nil(t, mes)
	assert.EqualError(t, err, "cannot have nil id")
}
