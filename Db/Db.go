package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func DbInit() {

	dsn := os.Getenv("DB_DSN")
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 👈 enable SQL log
	})

	if err != nil {
		log.Fatal("error occures when try to connect with database : ", err.Error())
	}
	DB = db
	sqlDB, _ := DB.DB()
	fmt.Println("→ Ping DB:", sqlDB.Ping()) // check connection
}
