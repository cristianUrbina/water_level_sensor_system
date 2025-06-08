package addsensmeastestutils

import (
	"bytes"
	"cristianUrbina/water_level_sensor_system/internal/dto"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func BuildRequest(sensorID uuid.UUID, meas *dto.AddSensorMeasurementHTTPRequestBody) (*http.Request, error) {
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

// func RunRequest(apiHandler *api.AddSensorMeasurementAPIHandler, req *http.Request) *httptest.ResponseRecorder {
// 	w := httptest.NewRecorder()
// 	apiHandler.ServeHTTP(w, req)
// 	return w
// }
