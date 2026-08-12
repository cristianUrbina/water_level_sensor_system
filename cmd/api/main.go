package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cristianUrbina/water_level_sensor_system/internal/api"
	"cristianUrbina/water_level_sensor_system/internal/application"
	"cristianUrbina/water_level_sensor_system/internal/infrastructure/mqtt"
	"cristianUrbina/water_level_sensor_system/internal/infrastructure/persistence"
	"cristianUrbina/water_level_sensor_system/internal/infrastructure/persistence/mysqlsensormeasurement"

	dbutils "cristianUrbina/water_level_sensor_system/pkg"

	sensormeasurementapp "cristianUrbina/water_level_sensor_system/internal/application/sensor_measurement"
	mysqlsensor "cristianUrbina/water_level_sensor_system/internal/infrastructure/persistence/mysql/sensor"

	"github.com/gorilla/mux"
	"github.com/mehdihadeli/go-mediatr"
)

func main() {
	// dsn := "cristian:cris2001@tcp(localhost:3306)/home_iot?parseTime=true"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db := dbutils.ConnectWithRetry(dsn)
	// if err != nil {
	// 	log.Fatalf("failed to connect to DB: %v", err)
	// }
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
	_, err = mqtt.NewPahoClient(mqtt.Config{
		BrokerURL: "tcp://192.168.1.125:1883",
		ClientID:  fmt.Sprintf("water-level-api-%d", time.Now().Unix()),
		Username:  "mqttuser",
		Password:  "cris2001",
	})
	if err != nil {
		log.Fatal(err)
	}

	readingsRepo, err := persistence.NewMySQLReadingsRepository(db)
	if err != nil {
		log.Fatalf("error creating readings repo: %v", err)
	}
	sensorDomainRepo, err := persistence.NewMySQLSensorRepository(db)
	if err != nil {
		log.Fatalf("error creating readings repo: %v", err)
	}
	addReadingHandler := application.NewAddSensorReadingHandler(readingsRepo, sensorDomainRepo)
	err = mediatr.RegisterRequestHandler(addReadingHandler)
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
