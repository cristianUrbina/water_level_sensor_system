package basedb

import (
	"cristianUrbina/water_level_sensor_system/internal/domain"
	"cristianUrbina/water_level_sensor_system/testutils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SetInitialData(db *gorm.DB) {
	ID, _ := uuid.Parse("002dcf46-2904-4478-9868-a7ff3f2bf1fc")
	// sensors := []*sensordm.Sensor{
	// 	&sensordm.Sensor{
	// 		ID: ID,
	// 		Name: "tanksensor",
	// 		Description: "sensor for checking tank water level",
	// 	},
	// }
	// testutils.AddSensors(db, sensors)
	sensors := []*domain.Sensor{
	{
		ID: ID,
		Profile: domain.ProfileTankLevel,
		Capabilities: []domain.Capability{
			domain.CapabilityDistance,
		},
		Metadata: domain.SensorMetadata{
			Name:            "Tank Sensor 1",
			Description:     "Ultrasonic water tank sensor",
			HardwareModel:   "JSN-SR04T",
		},
	},
	}
	testutils.AddDomainSensors(db, sensors)
}
