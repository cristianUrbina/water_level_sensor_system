package mysqlsensormeasurement

import (
	"log"
	"testing"

	sensormeasurement "cristianUrbina/water_level_sensor_system/internal/domain/sensor_measurement"
	"cristianUrbina/water_level_sensor_system/internal/domain/sensordm"
	"cristianUrbina/water_level_sensor_system/testutils"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	var cleanup func()
	var err error
	db, cleanup, err = testutils.CreateDockerizedMySQLDB()
	if err != nil {
		log.Fatalf("error creating dockerized db %v", err)
	}
	defer cleanup()
	m.Run()
}

func TestAddSensorMeasurement(t *testing.T) {
	sens, err := testutils.NewSensorBuilder().WithDefaultValues().Build()
	if err != nil {
		t.Fatalf("error building sensor: %v", err)
	}
	testutils.AddSensors(db, []*sensordm.Sensor{sens})
	mes, err := testutils.NewSensorMeasurmentBuilder().WithDefaultValues().WithSensorID(sens.ID).Build()
	if err != nil {
		t.Fatalf("error constructing sensor measurement: %v", err)
	}
	repo, err := NewMySQLSensorMeasurementRepository(db)
	if err != nil {
		t.Fatalf("error creating repository: %v", err)
	}
	err = repo.AddSensoreMeasurement(mes);
	assert.Nil(t, err)
	meas := []sensormeasurement.SensorMeasurement{}
	db.Find(&meas)
	assert.NotEmpty(t, meas)
	assert.Equal(t, mes.SensorID, meas[0].SensorID)
}
