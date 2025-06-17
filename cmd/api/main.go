package main

import (
	"log"
	"net/http"

	"cristianUrbina/water_level_sensor_system/internal/api"
	"cristianUrbina/water_level_sensor_system/internal/infrastructure/persistence/mysqlsensormeasurement"

	sensormeasurementapp "cristianUrbina/water_level_sensor_system/internal/application/sensor_measurement"
	mysqlsensor "cristianUrbina/water_level_sensor_system/internal/infrastructure/persistence/mysql/sensor"

	"github.com/gorilla/mux"
	"github.com/mehdihadeli/go-mediatr"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "cristian:cris2001@tcp(localhost:3306)/home_iot?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	sensorRepo, err := mysqlsensor.NewMySQLSensorRepository(db)
	if err != nil {
		log.Fatalf("failed to create repo: %v", err)
	}

	sensorMeasRepo, err := mysqlsensormeasurement.NewMySQLSensorMeasurementRepository(db)
	if err != nil {
		log.Fatalf("failed to create repo: %v", err)
	}

	appHandler := sensormeasurementapp.NewAddSensorMeasurementHandler(sensorMeasRepo, sensorRepo)
	err = mediatr.RegisterRequestHandler(appHandler)
	if err != nil {
		log.Fatalf("failed to register handler: %v", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/sensor/{sensorID}/measurement", api.NewAddSensorMeasurementAPIHandler().ServeHTTP).Methods("POST")
	log.Println("Starting server on :8080")
	if err = http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
