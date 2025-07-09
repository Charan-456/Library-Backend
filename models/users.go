package models

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"user_name" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
}

type Creds struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

var UserList = make(map[string]User)

var DB *gorm.DB

func DumpData() {
	Admin := User{
		Username: "admin",
		Password: "admin123",
		Email:    "admin@gmail.com",
	}
	UserList[Admin.Username] = Admin
}

func ConnectDB() {
	dbUsername := os.Getenv("dbUser")
	dbPassword := os.Getenv("dbPassword")
	dbName := os.Getenv("dbName")
	dbAddress := os.Getenv("hostAddress")
	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbAddress + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	dataBase, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting DB, Check the host address and port,password,user")
	}
	DB = dataBase
	log.Println("DB Connection set")
	DB.AutoMigrate(&User{})
	log.Println("Created Tabless ")
	admin := User{Username: "Admin", Password: "Admin1234", Email: "Admin@admin"}
	DB.FirstOrCreate(&admin)
}
