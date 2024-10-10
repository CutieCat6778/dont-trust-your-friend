package database

import (
	"cutiecat6778/dont-trust-your-friend/lib"
	"cutiecat6778/dont-trust-your-friend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(lib.POSTGES_URI), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})

	return db
}
