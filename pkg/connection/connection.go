package connection

import (
	"book-crud/pkg/config"
	"book-crud/pkg/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// connect database
func Connect() {
	dbConfig := config.LocalConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPass, dbConfig.DBIp, dbConfig.DBName)
	//root:nisaad@tcp(127.0.0.1:3306)/go-crud?charset=utf8mb4&parseTime=True&loc=Local
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("Error connecting to DB")
		panic(err)
	}
	fmt.Println("Database Connected")
	db = d
}

// creating new table in bookstore database
func migrate() {
	db.Migrator().AutoMigrate(&models.UserDetail{})
	db.Migrator().AutoMigrate(&models.BookDetail{})
}

// calling the connect function to initialize connection
func GetDB() *gorm.DB {
	if db == nil {
		Connect()
	}
	migrate()
	return db
}
