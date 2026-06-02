package persistence

import (
	"context"
	"errors"

	"cristianUrbina/water_level_sensor_system/internal/domain"

	"gorm.io/gorm"
)

func NewMySQLReadingsRepository(db *gorm.DB) (*MySQLSensorReadingsRepository, error) {
	if db == nil {
		return nil, errors.New("db cannot be nil")
	}
	return &MySQLSensorReadingsRepository{
		db: db,
	}, nil
}

type MySQLSensorReadingsRepository struct {
	db *gorm.DB
}

func (m *MySQLSensorReadingsRepository) Add(ctx context.Context, reading *domain.SensorReading) error {
	return m.db.Create(reading).Error
}
