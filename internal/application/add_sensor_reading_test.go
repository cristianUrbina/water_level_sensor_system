package application_test

import (
	"context"
	"testing"
	"time"

	"cristianUrbina/water_level_sensor_system/internal/application"
	appmocks "cristianUrbina/water_level_sensor_system/internal/application/mocks"
	"cristianUrbina/water_level_sensor_system/internal/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAddSensorReadingWithValidReading(t *testing.T) {
	ctrl := gomock.NewController(t)
	readings := appmocks.NewMockIReadingRepository(ctrl)
	sensors := appmocks.NewMockISensorRepository(ctrl)
	sensor := domain.Sensor{
		ID: uuid.New(),
	}

	sensors.EXPECT().GetByID(gomock.Any(), sensor.ID).Return(sensor, nil)
	readings.EXPECT().Add(gomock.Any(), gomock.Cond(func(x domain.SensorReading) bool { return x.SensorID == sensor.ID })).Return(nil)
	handler := application.NewAddSensorReadingHandler(readings, sensors)
	query := application.AddSensorReadingQuery{
		SensorID:   sensor.ID,
		Capability: "distance",
		Timestamp:  time.Now(),
		Value:      30,
	}
	err := handler.Handle(context.Background(), query)
	assert.NoError(t, err)
}

func TestAddSensorReadingForUnexistentSensorID(t *testing.T) {
	ctrl := gomock.NewController(t)
	readings := appmocks.NewMockIReadingRepository(ctrl)
	sensors := appmocks.NewMockISensorRepository(ctrl)
	sensor := domain.Sensor{
		ID: uuid.New(),
	}

	sensors.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(domain.Sensor{}, application.ErrRecordNotFound)
	handler := application.NewAddSensorReadingHandler(readings, sensors)
	query := application.AddSensorReadingQuery{
		SensorID:   sensor.ID,
		Capability: "distance",
		Timestamp:  time.Now(),
	}
	err := handler.Handle(context.Background(), query)
	assert.ErrorIs(t, application.ErrSensorNotFound, err)
}

func TestAddSensorReadingBadReading(t *testing.T) {
	ctrl := gomock.NewController(t)
	readings := appmocks.NewMockIReadingRepository(ctrl)
	sensors := appmocks.NewMockISensorRepository(ctrl)
	sensor := domain.Sensor{
		ID: uuid.New(),
	}

	sensors.EXPECT().GetByID(gomock.Any(), sensor.ID).Return(sensor, nil)
	handler := application.NewAddSensorReadingHandler(readings, sensors)
	query := application.AddSensorReadingQuery{
		SensorID:   sensor.ID,
		Capability: "distance",
		Timestamp:  time.Time{},
		Value:      30,
	}
	err := handler.Handle(context.Background(), query)
	assert.ErrorIs(t, application.ErrInvalidEntity, err)
}
