package dto

import "time"

type AddSensorMeasurementHTTPRequestBody struct {
	MeasuredAt time.Time
	Value      float64
	Type       string
}
