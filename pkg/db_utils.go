package dbutils

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectWithRetry(dsn string) *gorm.DB {
	var db *gorm.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			return db
		}
		log.Printf("Waiting for DB... (%d/10): %v\n", i+1, err)
		time.Sleep(3 * time.Second)
	}
	log.Fatalf("Failed to connect to DB after retries: %v\n", err)
	return nil
}
