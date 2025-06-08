package sensormeasurementapp_test

import (
	"context"
	sensormeasurementapp "cristianUrbina/water_level_sensor_system/internal/application/sensor_measurement"
	"cristianUrbina/water_level_sensor_system/internal/common"
	sensormeasurement "cristianUrbina/water_level_sensor_system/internal/domain/sensor_measurement"
	"cristianUrbina/water_level_sensor_system/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAddSensorMeasurementHandler_ShouldReturnOk_WhenPassingGoodMeasurement(t *testing.T) {
	expectedMeas, err := testutils.NewSensorMeasurmentBuilder().WithDefaultValues().Build()
	if err != nil {
		t.Fatalf("error creating sensor measurment: %v", err)
	}
	query := &sensormeasurementapp.AddSensorMeasurementQuery{
		SensorID: expectedMeas.SensorID,
		MeasuredAt: expectedMeas.MeasuredAt,
		Value: expectedMeas.Value,
		Type: expectedMeas.Type,
	}

	ctrl := gomock.NewController(t)
	mockRepo := testutils.NewMockISensorMeasurementRepository(ctrl)

	mockRepo.EXPECT().AddSensoreMeasurement(
		gomock.Cond(func(sm *sensormeasurement.SensorMeasurement) bool {
			return sm.SensorID == expectedMeas.SensorID && sm.Value == expectedMeas.Value
		})).Return(nil)

	sensorRepoMock := testutils.NewMockISensorRepository(ctrl)
	sensor, err := testutils.NewSensorBuilder().WithID(query.SensorID).WithDefaultValues().Build()
	if err != nil {
		t.Fatalf("error constructing sensor: %v", err)
	}
	sensorRepoMock.EXPECT().GetById(query.SensorID).Return(sensor, nil)

	handler := sensormeasurementapp.NewAddSensorMeasurementHandler(mockRepo, sensorRepoMock)
	res, err := handler.Handle(context.Background(), query)

	assert.NotNil(t, res)
	assert.Nil(t, err)
}


func TestAddSensorMeasurementHandler_ShouldReturnError_WhenPassingMeasurementForNonExistingSensor(t *testing.T) {
	expectedMeas, err := testutils.NewSensorMeasurmentBuilder().WithDefaultValues().Build()
	if err != nil {
		t.Fatalf("error creating sensor measurment: %v", err)
	}
	query := &sensormeasurementapp.AddSensorMeasurementQuery{
		SensorID: expectedMeas.SensorID,
		MeasuredAt: expectedMeas.MeasuredAt,
		Value: expectedMeas.Value,
		Type: expectedMeas.Type,
	}

	ctrl := gomock.NewController(t)
	mockRepo := testutils.NewMockISensorMeasurementRepository(ctrl)
	sensorRepoMock := testutils.NewMockISensorRepository(ctrl)
	sensorRepoMock.EXPECT().GetById(query.SensorID).Return(nil, nil)

	handler := sensormeasurementapp.NewAddSensorMeasurementHandler(mockRepo, sensorRepoMock)
	res, err := handler.Handle(context.Background(), query)

	assert.Nil(t, res)
	assert.Contains(t, err.Error(), "Not found sensor")
	assert.IsType(t, &common.NotFoundSensor{}, err)
}
