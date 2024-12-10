package database

import (
	"cutiecat6778/dont-trust-your-friend/lib"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	myLogger logger.Interface
)

func ConnectToDB() (*gorm.DB, *lib.CustomError) {
	db, err := gorm.Open(postgres.Open(lib.POSTGES_URI), &gorm.Config{
		TranslateError: true,
		Logger:         myLogger,
	})

	if err != nil {
		return nil, &lib.CustomError{
			Code:    500,
			Message: "Failed to connect to database",
			By:      lib.DatabaseService,
		}
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// db.AutoMigrate(&models.User{})

	return db, nil
}
