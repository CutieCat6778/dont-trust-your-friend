package handlers

import (
	"cutiecat6778/dont-trust-your-friend/database"
	"cutiecat6778/dont-trust-your-friend/lib"
	"cutiecat6778/dont-trust-your-friend/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type DBHandler struct {
	*gorm.DB
}

func InitDB() *DBHandler {
	db := database.ConnectToDB()
	return &DBHandler{db}
}

func handleDBError(err error) *lib.CustomError {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return lib.NewError("Record not found", 404, lib.DatabaseService)
	} else if errors.Is(err, gorm.ErrDuplicatedKey) {
		return lib.NewError("Duplicated key", 400, lib.DatabaseService)
	} else if errors.Is(err, gorm.ErrInvalidData) {
		return lib.NewError("Invalid data", 400, lib.DatabaseService)
	} else if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return lib.NewError("Foreign key violated", 400, lib.DatabaseService)
	} else {
		fmt.Println(err)
		return lib.NewError(err.Error(), 500, lib.DatabaseService)
	}
}

func (h *DBHandler) GetDB() *gorm.DB {
	return h.DB
}

func (h *DBHandler) CloseDB() {
	sqlDB, _ := h.DB.DB()
	sqlDB.Close()
}

func (h *DBHandler) GetUserByID(id uint) (*models.User, *lib.CustomError) {
	var user models.User
	result := h.First(&user, id)
	if result.Error != nil {
		return nil, handleDBError(result.Error)
	}
	return &user, nil
}

func (h *DBHandler) GetUserByUsername(username string) (*models.User, *lib.CustomError) {
	var user models.User
	result := h.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, handleDBError(result.Error)
	}
	return &user, nil
}

func (h *DBHandler) CreateUser(user models.User) *lib.CustomError {
	tx := h.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	result := tx.Create(&user)
	if result.Error != nil {
		tx.Rollback()
		return handleDBError(result.Error)
	}

	return nil
}

func (h *DBHandler) UpdateUserBalance(username string, amount int) *lib.CustomError {
	user, err := h.GetUserByUsername(username)
	if err != nil {
		return err
	}
	user.Balance += amount
	result := h.Save(user)
	if result.Error != nil {
		return handleDBError(result.Error)
	}
	return nil
}
