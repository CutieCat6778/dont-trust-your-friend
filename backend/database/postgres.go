package database

import (
	"cutiecat6778/dont-trust-your-friend/lib"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, *lib.CustomError) {
	db, err := gorm.Open(postgres.Open(lib.POSTGES_URI), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return nil, &lib.CustomError{
			Code:    500,
			Message: "Failed to connect to database",
			By:      lib.DatabaseService,
		}
	}

	// db.AutoMigrate(&models.User{})

	return db, nil
}
