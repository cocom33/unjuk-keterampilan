package config

import (
	"project/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := "root@tcp(127.0.0.1:3306)/prakerja?charset=utf8mb4&parseTime=True&loc=Local"
  var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("gagal init database")
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(&models.Student{})
}

func GetDBInstance() *gorm.DB {
	return DB
}