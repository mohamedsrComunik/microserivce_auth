package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mohamedsrComunik/microserivce_auth/model"
)

// Connexion to connect to the db
func Connexion() *gorm.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")
	host := "mysql_service"
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, database))
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&model.User{})
	return db
}
