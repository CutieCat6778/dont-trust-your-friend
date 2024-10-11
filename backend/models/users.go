package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type User struct {
	Model
	Username string `json:"username" gorm:"unique;not null;size:255"`
	Password []byte `json:"password" gorm:"not null;size:64"`
	Name     string `json:"name"`
	Balance  int    `json:"balance"`
	Version  int    `json:"version" gorm:"not null,default:0"`
}
