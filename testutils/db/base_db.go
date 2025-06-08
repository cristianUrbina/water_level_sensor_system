package basedb

import (
	"cristianUrbina/water_level_sensor_system/internal/domain/sensordm"
	"cristianUrbina/water_level_sensor_system/testutils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SetInitialData(db *gorm.DB) {
	ID,_ := uuid.Parse("8b7f1c5a-3e8f-47cc-a7bc-b8610d489b56")
	sensors := []*sensordm.Sensor{
		&sensordm.Sensor{
			ID: ID,
			Name: "tanksensor",
			Description: "sensor for checking tank water level",
		},
	}
	testutils.AddSensors(db, sensors)
}
