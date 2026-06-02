package mqtt_test

import (
	"context"
	"log"
	"testing"
	"time"

	"cristianUrbina/water_level_sensor_system/internal/application"
	"cristianUrbina/water_level_sensor_system/internal/infrastructure/mqtt"
	pkg "cristianUrbina/water_level_sensor_system/pkg/errors"

	"github.com/google/uuid"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/stretchr/testify/assert"
)

type MediatorHandler struct {
	Received *application.AddSensorReadingQuery
}

func (m *MediatorHandler) Handle(ctx context.Context, query application.AddSensorReadingQuery) (mediatr.Unit, error) {
	log.Println("query", query)
	m.Received = &query
	return mediatr.Unit{}, nil
}

func (m MediatorHandler) GetReceived() *application.AddSensorReadingQuery {
	return m.Received
}

func TestReadingsHandlerHandle(t *testing.T) {
	msg := mqtt.Message{
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
	// ctrl := gomock.NewController(t)
	mediatorHandler := &MediatorHandler{}
	// mediator := appmocks.NewMockMediator(ctrl)
	// mediator.EXPECT().Send(gomock.Any(), expectedQuery).Return(nil)
	mediatr.ClearRequestRegistrations()
	err := mediatr.RegisterRequestHandler[application.AddSensorReadingQuery](mediatorHandler)
	if err != nil {
		t.Fatalf("failed to register handler: %v", err)
	}
	handler := mqtt.NewReadingHandler()
	error := handler.Handle(msg)
	assert.NoError(t, error)
	assert.Equal(t, expectedQuery.Value, mediatorHandler.GetReceived().Value)
}

func TestReadingsHandlerHandleBadSensorID(t *testing.T) {
	q := mqtt.Message{
		Topic: "sensor01",
		Payload: []byte(`{
        "sensor_id": "invalid-uuid",
        "capability": "temperature",
        "value": 23.5,
        "unit": "C",
        "timestamp": "2026-06-01T18:30:00Z"
        }`),
	}
	handler := mqtt.NewReadingHandler()
	error := handler.Handle(q)
	assert.ErrorIs(t, error, pkg.ErrInvalidUUID)
}

func TestReadingsHandlerHandleBadTimestamp(t *testing.T) {
	q := mqtt.Message{
		Topic: "sensor01",
		Payload: []byte(`{
        "sensor_id": "c56a4180-65aa-42ec-a945-5fd21dec0538",
        "capability": "temperature",
        "value": 23.5,
        "unit": "C",
        "timestamp": "not a timestamp"
        }`),
	}
	handler := mqtt.NewReadingHandler()
	error := handler.Handle(q)
	assert.ErrorIs(t, error, pkg.ErrInvalidTimestamp)
}
