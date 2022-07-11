package config

import (
	"chilindo/src/admin-service/entity"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

//SetupDatabaseConnection creating a new connection to database
func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load(".env.admin")
	if err != nil {
		panic("Failed to load env file")

	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&entity.Administrator{})
	return db
}

//CloseDatabaseConnection close database
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close database")
	}
	dbSQL.Close()
}
