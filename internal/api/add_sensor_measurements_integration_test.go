package api

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"cristianUrbina/water_level_sensor_system/internal/domain/sensordm"
	"cristianUrbina/water_level_sensor_system/internal/dto"
	"cristianUrbina/water_level_sensor_system/internal/infrastructure/persistence/mysqlsensormeasurement"
	"cristianUrbina/water_level_sensor_system/testutils"

	addsensmeastestutils "cristianUrbina/water_level_sensor_system/testutils/common"

	sensormeasurementapp "cristianUrbina/water_level_sensor_system/internal/application/sensor_measurement"

	mysqlsensor "cristianUrbina/water_level_sensor_system/internal/infrastructure/persistence/mysql/sensor"

	"github.com/mehdihadeli/go-mediatr"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB
	cleanup func()
)

func TestMain(m *testing.M) {
	var err error
	db, cleanup, err = testutils.CreateDockerizedMySQLDB()
	if err != nil {
		log.Fatalf("error creating dockerized db: %v", err)
	}
	code := m.Run()
	cleanup()
	os.Exit(code)
}

func TestAddSensorMeasurement_E2E_WithValidSensorMeasurement(t *testing.T) {
	mediatr.ClearRequestRegistrations()
	sensor, err := testutils.NewSensorBuilder().WithDefaultValues().Build()
	if err != nil {
		t.Fatalf("error building sensor: %v", err)
	}
	sensors := []*sensordm.Sensor{sensor}
	testutils.AddSensors(db, sensors)
	repo, err := mysqlsensormeasurement.NewMySQLSensorMeasurementRepository(db)
	if err != nil {
		t.Fatalf("error creating repo: %v", err)
	}

	sensorRepo, err := mysqlsensor.NewMySQLSensorRepository(db)
	if err != nil {
		t.Fatalf("error creating sensor repo: %v", err)
	}
	appHandler := sensormeasurementapp.NewAddSensorMeasurementHandler(repo, sensorRepo)
	mediatr.RegisterRequestHandler[*sensormeasurementapp.AddSensorMeasurementQuery, *sensormeasurementapp.AddSensorMeasurementResponse](appHandler)
	reqBody := &dto.AddSensorMeasurementHTTPRequestBody{
		MeasuredAt: time.Now(),
		Value:      34.2,
		Type:       "water_level_sensor",
	}
	req, err := addsensmeastestutils.BuildRequest(sensor.ID, reqBody)
	if err != nil {
		t.Fatalf("error building request: %v", err)
	}
	apiHandler := NewAddSensorMeasurementAPIHandler()

	w := httptest.NewRecorder()
	apiHandler.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
}
