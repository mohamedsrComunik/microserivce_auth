package model

import (
	"github.com/jinzhu/gorm"
)

// User struct
type User struct {
	gorm.Model
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
}
