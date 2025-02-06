package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email    string `json:"email" gorm:"uniqueIndex,size:256"`
	Password string `json:"password" gorm:"size:256"`
}
