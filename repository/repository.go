package repository

import (
	"fmt"

	"github.com/mohamedsrComunik/microserivce_auth/model"

	"github.com/jinzhu/gorm"
)

type userRepo struct {
	db *gorm.DB
}

// Repository interface
type Repository interface {
	Create(user *model.User) error
	GetUserByEmail(user *model.User) bool
}

// New to create new instance of userRepo with db
func New(db *gorm.DB) userRepo {
	u := userRepo{db}
	return u
}

// Create to add new user to the db
func (u *userRepo) Create(user *model.User) error {
	if err := u.db.Create(user).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (u *userRepo) GetUserByEmail(user *model.User) bool {
	if err := u.db.Model(model.User{}).Where("email = ?", user.Email).First(user).Error; err != nil {
		return false
	}
	return true
}
