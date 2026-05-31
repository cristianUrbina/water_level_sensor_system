package application_test

import (
	"context"
	"testing"
	"time"

	"cristianUrbina/water_level_sensor_system/internal/application"
	appmocks "cristianUrbina/water_level_sensor_system/internal/application/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandleSensorReadingMessage(t *testing.T) {
	q := &application.HandleSensorReadingMessageQuery{
		Topic: "sensor01",
		Payload: []byte(`{
        "sensor_id": "c56a4180-65aa-42ec-a945-5fd21dec0538",
        "capability": "temperature",
        "value": 23.5,
        "unit": "C",
        "timestamp": "2026-06-01T18:30:00Z"
        }`),
	}
	sID, _ := uuid.Parse("c56a4180-65aa-42ec-a945-5fd21dec0538")
	timestamp, _ := time.Parse(time.RFC3339, "2026-06-01T18:30:00Z")
	expectedQuery := application.AddSensorReadingQuery{
		SensorID:   sID,
		Capability: "temperature",
		Value:      23.5,
		Unit:       "C",
		Timestamp:  timestamp,
	}
	ctrl := gomock.NewController(t)
	mediator := appmocks.NewMockMediator(ctrl)
	mediator.EXPECT().Send(gomock.Any(), expectedQuery).Return(nil)
	handler := application.NewHandleSensorReadingMessageHandler(mediator)
	_, error := handler.Handle(context.Background(), q)
	assert.NoError(t, error)
}

func TestHandleSensorReadingMessageBadSensorID(t *testing.T) {
	q := &application.HandleSensorReadingMessageQuery{
		Topic: "sensor01",
		Payload: []byte(`{
        "sensor_id": "invalid-uuid",
        "capability": "temperature",
        "value": 23.5,
        "unit": "C",
        "timestamp": "2026-06-01T18:30:00Z"
        }`),
	}
	ctrl := gomock.NewController(t)
	mediator := appmocks.NewMockMediator(ctrl)
	handler := application.NewHandleSensorReadingMessageHandler(mediator)
	_, error := handler.Handle(context.Background(), q)
	assert.ErrorIs(t, error, application.ErrInvalidUUID)
}
