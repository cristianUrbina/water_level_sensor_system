package sensordm

import (
	"cristianUrbina/water_level_sensor_system/internal/domain/tank"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestWaterLevelSensorConstructor(t *testing.T) {
	tank :=  tank.NewTank()
	sens, err := GetWaterLevelSensor(uuid.New(), "testing_sensor", "sensor used for testing", tank)
	assert.Nil(t, err)
	assert.NotNil(t, sens)
}

func TestWaterLevelSensorConstructorShouldReturnErrorWhenPassingNilTankShould(t *testing.T) {
	sens, err := GetWaterLevelSensor(uuid.New(), "testing_sensor", "sensor used for testing", nil)
	assert.Error(t, err)
	assert.Nil(t, sens)
}
