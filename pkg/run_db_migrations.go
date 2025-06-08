package main

import (
	"cristianUrbina/water_level_sensor_system/testutils"
	basedb "cristianUrbina/water_level_sensor_system/testutils/db"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "cristian:cris2001@tcp(localhost:3306)/home_iot?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	log.Println("Connected to DB, running migrations")
	err = testutils.RunMigrations(db)
	if err != nil {
		log.Fatalf("error runnning migrations: %v", err)
	}
	log.Println("Migration complete")
	basedb.SetInitialData(db)
}
