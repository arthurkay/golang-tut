package db

import (
	"os"

	"db_proj/models"

	"github.com/arthurkay/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() (*gorm.DB, error) {
	env.Load("../.env")

	dbname := os.Getenv("DBNAME")
	host := os.Getenv("HOST")
	user := os.Getenv("USERNAME")
	pass := os.Getenv("PASSWORD")
	mysqlPort := os.Getenv("MYSQL_PORT")

	dsn := user + ":" + pass + "@tcp(" + host + ":" + mysqlPort + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return db, err
	}

	db.AutoMigrate(&models.User{}, &models.Address{})
	return db, nil
}

func CreateUser(db *gorm.DB, user *models.User) {
	db.Create(user)
}

func GetFirstUser(db *gorm.DB) *models.User {
	user := models.User{}
	db.First(&user)
	return &user
}
