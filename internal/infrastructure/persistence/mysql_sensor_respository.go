package persistence

import (
	"context"
	"errors"

	"cristianUrbina/water_level_sensor_system/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewMySQLSensorRepository(db *gorm.DB) (*MySQLSensorRepository, error) {
	if db == nil {
		return nil, errors.New("db cannot be nil")
	}
	return &MySQLSensorRepository{
		db: db,
	}, nil
}

type MySQLSensorRepository struct {
	db *gorm.DB
}

func (m *MySQLSensorRepository) GetByID(
	ctx context.Context,
	ID uuid.UUID,
) (domain.Sensor, error) {
	var sensor domain.Sensor

	err := m.db.WithContext(ctx).
		First(&sensor, "id = ?", ID.String()).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Sensor{}, err
		}
		return domain.Sensor{}, err
	}

	return sensor, nil
}
