package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	sensormeasurementapp "cristianUrbina/water_level_sensor_system/internal/application/sensor_measurement"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mehdihadeli/go-mediatr"
)

func NewAddSensorMeasurementAPIHandler() *AddSensorMeasurementAPIHandler {
	return &AddSensorMeasurementAPIHandler{}
}

type AddSensorMeasurementHTTPRequestBody struct {
	MeasuredAt time.Time
	Value      float64
	Type       string
}

type AddSensorMeasurementAPIHandler struct{}

func (a *AddSensorMeasurementAPIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sensIDStr := vars["sensorID"]
	sensID, err := uuid.Parse(sensIDStr)
	if err != nil {
		http.Error(w, "invalid sensor ID: " + err.Error(), http.StatusBadRequest)
	}

	var reqBody AddSensorMeasurementHTTPRequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "invalid JSON body: " + err.Error(), http.StatusBadRequest)
	}
	query := &sensormeasurementapp.AddSensorMeasurementQuery{
		SensorID: sensID,
		MeasuredAt: reqBody.MeasuredAt,
		Value: reqBody.Value,
		Type: reqBody.Type,
	}
	_, err = mediatr.Send[*sensormeasurementapp.AddSensorMeasurementQuery, *sensormeasurementapp.AddSensorMeasurementResponse](context.Background(), query)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
