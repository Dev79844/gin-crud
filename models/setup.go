package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB



func ConnectDB(){
	// dsn := getURI("MYSQL_URI")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
					os.Getenv("MYSQL_USER"),
					os.Getenv("MYSQL_PASS"),
					os.Getenv("MYSQL_HOST"),
					os.Getenv("MYSQL_DB"),
				)
	// dsn := "dev:dev1234@tcp(127.0.0.1:3306)/recordings?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(db)
	fmt.Println("Connected")

	db.AutoMigrate(&Post{})

	DB = db
}
