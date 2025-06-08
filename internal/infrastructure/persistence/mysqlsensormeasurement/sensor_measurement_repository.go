package mysqlsensormeasurement

import (
	sensormeasurement "cristianUrbina/water_level_sensor_system/internal/domain/sensor_measurement"
	"errors"

	"gorm.io/gorm"
)

func NewMySQLSensorMeasurementRepository(db *gorm.DB) (*MySQLSensorMeasurementRepository, error) {
	if db == nil {
		return nil, errors.New("db cannot be nil")
	}
	return &MySQLSensorMeasurementRepository{
		db: db,
	}, nil
}

type MySQLSensorMeasurementRepository struct {
	db *gorm.DB
}

func (m *MySQLSensorMeasurementRepository) AddSensoreMeasurement(meas *sensormeasurement.SensorMeasurement) error {
	return m.db.Create(meas).Error
}
