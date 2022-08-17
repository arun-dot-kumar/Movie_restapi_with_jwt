package database

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"movie/models"
)

var Instance *gorm.DB
var dbError error

func Connect(connstr string) {
	Instance, dbError = gorm.Open(mysql.Open(connstr), &gorm.Config{})
	//DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbUser"), os.Getenv("TestDbName"), os.Getenv("TestDbPassword"))
	//Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	//conn, err := sql.Open("mysql", "root:mypassword@tcp(db:3306)/testdb")

	if dbError != nil {
		fmt.Println(dbError)
		panic("connection error")
	}
	log.Println("Connected to Database!")
}

func Migrate() {
	err := Instance.AutoMigrate(&models.User{}, &models.AdminUser{}, &models.Show{}, &models.BookedList{})
	if err != nil {
		return
	}
	log.Println("Database Migration Completed!")
}
