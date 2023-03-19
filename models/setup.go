package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"os"
	"log"
)

var DB *gorm.DB

func getURI(key string) string{
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	
	return os.Getenv(key)
}

func ConnectDB(){
	dsn := getURI("MYSQL_URI")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		panic("Failed to connect to database")
	}
	fmt.Println(db)
	fmt.Println("Connected")

	db.AutoMigrate(&Post{})

	DB = db
}
