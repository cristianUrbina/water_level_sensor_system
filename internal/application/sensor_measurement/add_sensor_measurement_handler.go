package sensormeasurementapp

import (
	"context"

	"cristianUrbina/water_level_sensor_system/internal/common"
	sensormeasurement "cristianUrbina/water_level_sensor_system/internal/domain/sensor_measurement"
	"cristianUrbina/water_level_sensor_system/internal/domain/sensordm"

	"github.com/google/uuid"
)

type IAddSensorMeasurementHandler interface {
	Handle(ctx context.Context, query *AddSensorMeasurementQuery) (*AddSensorMeasurementResponse, error)
}

func NewAddSensorMeasurementHandler(repo sensormeasurement.ISensorMeasurementRepository, sensorRepo sensordm.ISensorRepository) *AddSensorMeasurementHandler {
	return &AddSensorMeasurementHandler{
		repo: repo,
		sensorRepo: sensorRepo,
	}
}

type AddSensorMeasurementHandler struct {
	repo sensormeasurement.ISensorMeasurementRepository
	sensorRepo sensordm.ISensorRepository
}

func (a *AddSensorMeasurementHandler) Handle(ctx context.Context, query *AddSensorMeasurementQuery) (*AddSensorMeasurementResponse, error) {
	sens, err := a.sensorRepo.GetById(query.SensorID)
	if err != nil || sens == nil {
		return nil, common.NewNotFoundSensorError(query.SensorID)
	}
	sensMeas, err := sensormeasurement.NewSensorMeasurement(uuid.New(), sens.ID, query.MeasuredAt, query.Value, query.Type)
	if err != nil {
		return nil, err
	}

	err = a.repo.AddSensoreMeasurement(sensMeas)
	if err != nil {
		return nil, err
	}

	return &AddSensorMeasurementResponse{}, nil
}
