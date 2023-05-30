package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDB(){
	dsn := fmt.Sprintf("%s:%s@tcp(host.docker.internal:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",os.Getenv("MYSQL_USER"),os.Getenv("MYSQL_PASS"),os.Getenv("MYSQL_DB"))
	// dsn := getURI("MYSQL_URI")
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		panic("Failed to connect to database")
	}
	// fmt.Println(db)
	fmt.Println("Connected")

	db.AutoMigrate(&Post{})

	DB = db
}
