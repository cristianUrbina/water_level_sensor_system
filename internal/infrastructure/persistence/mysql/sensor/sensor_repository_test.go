package mysqlsensor

import (
	"log"
	"testing"

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

func TestGetById(t *testing.T) {
	sens, err := testutils.NewSensorBuilder().WithDefaultValues().Build()
	if err != nil {
		t.Fatalf("error constructing sensor: %v", err)
	}
	testutils.AddSensors(db, []*sensordm.Sensor{sens})
	repo, err := NewMySQLSensorRepository(db)
	if err != nil {
		t.Fatalf("error creating repository: %v", err)
	}
	res, err := repo.GetById(sens.ID);
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, sens, res)
}
