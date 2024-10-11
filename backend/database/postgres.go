package database

import (
	"cutiecat6778/dont-trust-your-friend/lib"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(lib.POSTGES_URI), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(&models.User{})

	return db
}
