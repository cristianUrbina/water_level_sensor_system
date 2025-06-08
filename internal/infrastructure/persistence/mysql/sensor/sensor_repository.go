package mysqlsensor

import (
	"cristianUrbina/water_level_sensor_system/internal/domain/sensordm"

	"github.com/google/uuid"
	"gorm.io/gorm"

	dal "cristianUrbina/water_level_sensor_system/internal/infrastructure/persistence/mysql/query"
)

func NewMySQLSensorRepository(db *gorm.DB) (*MySQLSensorRepository, error) {
	return &MySQLSensorRepository{
		db: db,
	}, nil
}

type MySQLSensorRepository struct {
	db *gorm.DB
}


func (m *MySQLSensorRepository) GetById(id uuid.UUID) (*sensordm.Sensor, error) {
	dal.SetDefault(m.db)
	sensor, err := dal.Q.Sensor.Where(dal.Sensor.ID.Eq(id)).First()
	return sensor, err
}
