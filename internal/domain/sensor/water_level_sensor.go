package sensor

import (
	"cristianUrbina/water_level_sensor_system/internal/domain/tank"

	"github.com/google/uuid"
)

func GetWaterLevelSensor(id uuid.UUID, name, description string, tank *tank.Tank) *WaterLevelSensor {
	return &WaterLevelSensor{
		id: id,
		name: name,
		description: description,
		tank: tank,
	}
}

type Sensor interface {
	GetLecture()
}

type WaterLevelSensor struct {
	id uuid.UUID
	name string
	description string
	tank *tank.Tank
}

