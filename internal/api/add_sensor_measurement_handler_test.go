package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"cristianUrbina/water_level_sensor_system/internal/common"
	"cristianUrbina/water_level_sensor_system/internal/dto"
	"cristianUrbina/water_level_sensor_system/testutils"
	addsensmeastestutils "cristianUrbina/water_level_sensor_system/testutils/common"

	sensmesapp "cristianUrbina/water_level_sensor_system/internal/application/sensor_measurement"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func setup(t *testing.T) *testutils.MockIAddSensorMeasurementHandler {
	t.Helper()
	mediatr.ClearRequestRegistrations()

	ctrl := gomock.NewController(t)
	return testutils.NewMockIAddSensorMeasurementHandler(ctrl)
}

func newAppQuery(sensorID uuid.UUID, reqBody *dto.AddSensorMeasurementHTTPRequestBody) *sensmesapp.AddSensorMeasurementQuery {
	return &sensmesapp.AddSensorMeasurementQuery{
		SensorID:   sensorID,
		MeasuredAt: reqBody.MeasuredAt.Round(0),
		Value:      reqBody.Value,
		Type:       reqBody.Type,
	}
}

func buildRequest(sensorID uuid.UUID, meas *dto.AddSensorMeasurementHTTPRequestBody) (*http.Request, error) {
	url := fmt.Sprintf("/sensor/%s/measurement", sensorID)
	jsonBody, err := json.Marshal(meas)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(jsonBody)
	req := httptest.NewRequest(http.MethodPost, url, body)
	req = mux.SetURLVars(req, map[string]string{
		"sensorID": sensorID.String(),
	})
	return req, nil
}

func runRequest(apiHandler *AddSensorMeasurementAPIHandler, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	apiHandler.ServeHTTP(w, req)
	return w
}

func TestServeHTTPWithValidMeasurement(t *testing.T) {
	handlerMock := setup(t)
	reqBody := &dto.AddSensorMeasurementHTTPRequestBody{
		MeasuredAt: time.Now(),
		Value:      32.4,
		Type:       "water_level",
	}
	sensorID := uuid.New()
	query := newAppQuery(sensorID, reqBody)

	handlerMock.EXPECT().Handle(context.Background(), query).Return(nil, nil)
	mediatr.RegisterRequestHandler[*sensmesapp.AddSensorMeasurementQuery, *sensmesapp.AddSensorMeasurementResponse](handlerMock)

	req, err := addsensmeastestutils.BuildRequest(sensorID, reqBody)
	if err != nil {
		t.Fatalf("error building request: %v", err)
	}
	apiHandler := NewAddSensorMeasurementAPIHandler()

	w := runRequest(apiHandler, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestServeHTTPWithNonExistingSensor(t *testing.T) {
	handlerMock := setup(t)
	reqBody := &dto.AddSensorMeasurementHTTPRequestBody{
		MeasuredAt: time.Now(),
		Value:      32.4,
		Type:       "water_level",
	}
	sensorID := uuid.New()
	query := newAppQuery(sensorID, reqBody)

	expectedError := common.NewNotFoundSensorError(sensorID)
	handlerMock.EXPECT().Handle(context.Background(), query).Return(nil, expectedError)
	mediatr.RegisterRequestHandler[*sensmesapp.AddSensorMeasurementQuery, *sensmesapp.AddSensorMeasurementResponse](handlerMock)

	req, err := addsensmeastestutils.BuildRequest(sensorID, reqBody)
	if err != nil {
		t.Fatalf("error building request: %v", err)
	}

	apiHandler := NewAddSensorMeasurementAPIHandler()

	w := runRequest(apiHandler, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
