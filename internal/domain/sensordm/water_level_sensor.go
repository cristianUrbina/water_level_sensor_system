package sensordm

import (
	"errors"

	sensormeasurement "cristianUrbina/water_level_sensor_system/internal/domain/sensor_measurement"
	"cristianUrbina/water_level_sensor_system/internal/domain/tank"

	"github.com/google/uuid"
)

func GetWaterLevelSensor(id uuid.UUID, name, description string, tank *tank.Tank) (*Sensor, error) {
	if tank == nil {
		return nil, errors.New("tank cannot be nil")
	}

	return &Sensor{
		ID:          id,
		Name:        name,
		Description: description,
		// Tank: tank,
	}, nil
}

// type Sensor interface {
// 	GetLecture()
// }

type Sensor struct {
	ID           uuid.UUID `gorm:"type:varchar(36);primaryKey"`
	Name         string
	Description  string
	Measurements []sensormeasurement.SensorMeasurement `gorm:"foreignKey:SensorID"`
	// TankId uuid.UUID `gorm:1
	// Tank *tank.Tank
}
