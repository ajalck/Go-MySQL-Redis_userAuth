package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"user_name"`
	Email    string `json:"email" gorm:"unique" binding:"required,email"`
	Password string `json:"password" gorm:"unique" binding:"required,min=6"`
}
