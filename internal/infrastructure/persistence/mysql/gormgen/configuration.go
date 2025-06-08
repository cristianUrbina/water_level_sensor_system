package main

import (
	"cristianUrbina/water_level_sensor_system/internal/domain/sensordm"
	sensormeasurement "cristianUrbina/water_level_sensor_system/internal/domain/sensor_measurement"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/infrastructure/persistence/mysql/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	// Use a dummy in-memory DB just for syntax validation if needed
	dsn := "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Proceeding without live DB (optional)")
	}
	g.UseDB(db)

	// Use your manually defined models
	g.ApplyBasic(sensordm.Sensor{})
	g.ApplyBasic(sensormeasurement.SensorMeasurement{})

	g.Execute()
}
