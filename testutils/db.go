package testutils

import (
	"cristianUrbina/water_level_sensor_system/internal/domain/sensordm"
	sensormeasurement "cristianUrbina/water_level_sensor_system/internal/domain/sensor_measurement"
	"errors"
	"fmt"
	"log"

	"github.com/ory/dockertest/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateDockerizedMySQLDB() (*gorm.DB, func(), error) {
	var db *gorm.DB
	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("Could not construct pool: %s", err))
	}

	err = pool.Client.Ping()
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("Could not connect to Docker: %s", err))
	}

	resource, err := pool.Run("mysql", "8.0.42", []string{"MYSQL_ROOT_PASSWORD=secret"})
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("Could not start resource: %s", err))
	}

	if err := pool.Retry(func() error {
		var err error
		dsn := fmt.Sprintf("root:secret@tcp(localhost:%s)/mysql?parseTime=true", resource.GetPort("3306/tcp"))
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, nil, errors.New(fmt.Sprintf("Could not connect to database: %s", err))
	}
	cleanup := func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}

	err = RunMigrations(db)
	if err != nil {
		return nil, nil, err
	}

	return db, cleanup, nil
}

func RunMigrations(gdb *gorm.DB) error {
	err := gdb.AutoMigrate(
		&sensordm.Sensor{},
		&sensormeasurement.SensorMeasurement{},
		)
	return err
}

func AddSensors(gdb *gorm.DB, sensors []*sensordm.Sensor) {
	addObjects(gdb, sensors)
}

func AddSensorMeasurements(gdb *gorm.DB, meas []*sensormeasurement.SensorMeasurement){
	addObjects(gdb, meas)
}

func addObjects[T any](gdb *gorm.DB, objs []*T) {
	for _, s := range objs {
		gdb.Create(s)
	}
}
