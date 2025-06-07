package sensor

import (
	"cristianUrbina/water_level_sensor_system/internal/domain/tank"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestWaterLevelSensorConstructor(t *testing.T) {
	tank :=  tank.GetTank()
	sens := GetWaterLevelSensor(uuid.New(), "testing_sensor", "sensor used for testing", tank)
	assert.NotNil(t, sens)
}
