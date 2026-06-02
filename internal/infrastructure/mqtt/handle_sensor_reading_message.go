package mqtt

import (
	"context"
)

type IHandleSensorReadingMessageHandler interface {
	Handle(ctx context.Context, q HandleSensorReadingMessageQuery) error
}

type HandleSensorReadingMessageQuery struct {
	Topic   string
	Payload []byte
}

type sensorReadingDTO struct {
	SensorID   string  `json:"sensor_id"`
	Capability string  `json:"capability"`
	Value      float64 `json:"value"`
	Unit       string  `json:"unit"`
	Timestamp  string  `json:"timestamp"`
}

// func NewHandleSensorReadingMessageHandler(m Mediator) HandleSensorReadingMessageHandler {
// 	return HandleSensorReadingMessageHandler{
// 		mediator: m,
// 	}
// }
//
// type HandleSensorReadingMessageHandler struct {
// 	mediator Mediator
// }

// func (h HandleSensorReadingMessageHandler) Handle(ctx context.Context, query *HandleSensorReadingMessageQuery) (mediatr.Unit, error) {
// 	var dto sensorReadingDTO
// 	if err := json.Unmarshal(query.Payload, &dto); err != nil {
// 		return mediatr.Unit{}, err
// 	}
// 	sensorID, err := uuid.Parse(dto.SensorID)
// 	if err != nil {
// 		return mediatr.Unit{}, ErrInvalidUUID
// 	}
// 	timestamp, err := time.Parse(time.RFC3339, "2026-06-01T18:30:00Z")
// 	if err != nil {
// 		return mediatr.Unit{}, err
// 	}
// 	q := AddSensorReadingQuery{
// 		SensorID:   sensorID,
// 		Capability: dto.Capability,
// 		Value:      dto.Value,
// 		Unit:       dto.Unit,
// 		Timestamp:  timestamp,
// 	}
// 	return mediatr.Unit{}, h.mediator.Send(ctx, q)
// }
