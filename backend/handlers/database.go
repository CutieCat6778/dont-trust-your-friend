package handlers

import (
	"cutiecat6778/dont-trust-your-friend/database"
	"cutiecat6778/dont-trust-your-friend/lib"
	"cutiecat6778/dont-trust-your-friend/models"
	"errors"

	"gorm.io/gorm"
)

type DBHandler struct {
	*gorm.DB
}

var (
	DB chan *DBHandler
)

func InitDB() chan *DBHandler {
	DB = make(chan *DBHandler)
	go func() {
		db := database.ConnectToDB()
		DB <- &DBHandler{db}
	}()
	return DB
}

func handleDBError(err error) *lib.CustomError {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return (lib.NewError("Record not found", 404, lib.DatabaseService))
	} else if errors.Is(err, gorm.ErrDuplicatedKey) {
		return (lib.NewError("Duplicated key", 400, lib.DatabaseService))
	} else {
		return (lib.NewError(err.Error(), 500, lib.DatabaseService))
	}
}

func (h *DBHandler) GetDB() *gorm.DB {
	return h.DB
}

func (h *DBHandler) GetUserByID(id uint) (*models.User, *lib.CustomError) {
	var user models.User
	result := h.First(&user, id)
	return &user, handleDBError(result.Error)
}

func (h *DBHandler) GetUserByUsername(username string) (*models.User, *lib.CustomError) {
	var user models.User
	result := h.Where("username = ?", username).First(&user)
	return &user, handleDBError(result.Error)
}

func (h *DBHandler) CreateUser(user *models.User) *lib.CustomError {
	result := h.Create(user)
	return handleDBError(result.Error)
}

func (h *DBHandler) UpdateUserBalance(username string, amount int) *lib.CustomError {
	user, err := h.GetUserByUsername(username)
	if err != nil {
		return err
	}
	user.Balance += amount
	result := h.Save(user)
	return handleDBError(result.Error)
}
