package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null;size:255"`
	Password []byte `json:"password" gorm:"not null;size:64"`
	Name     string `json:"name"`
	Balance  int    `json:"balance"`
}
