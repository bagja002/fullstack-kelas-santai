package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"backend1/models"

)

var DB *gorm.DB

func Connect(){

	dsn := "root@tcp(127.0.0.1:3306)/projek2?charset=utf8mb4&parseTime=True&loc=Local"
	con, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {	
		fmt.Println("Database tidak dapat terhubung")
	}
	fmt.Println("database terhubung")
	DB = con
	DB.AutoMigrate(models.Admin{}, models.User{})
}