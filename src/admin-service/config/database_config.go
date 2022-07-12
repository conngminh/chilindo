package config

import (
	"chilindo/src/admin-service/entity"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

////SetupDatabaseConnection creating a new connection to database
//func SetupDatabaseConnection() *gorm.DB {
//	err := godotenv.Load(".env.admin")
//	if err != nil {
//		panic("Failed to load env file")
//
//	}
//	dbUser := os.Getenv("DB_USER")
//	dbPass := os.Getenv("DB_PASS")
//	dbHost := os.Getenv("DB_HOST")
//	dbName := os.Getenv("DB_NAME")
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic("Failed to connect database")
//	}
//	db.AutoMigrate(&entity.Administrator{})
//	return db
//}
//
////CloseDatabaseConnection close database
//func CloseDatabaseConnection(db *gorm.DB) {
//	dbSQL, err := db.DB()
//	if err != nil {
//		panic("Failed to close database")
//	}
//	dbSQL.Close()
//}
var (
	host     string = "localhost"
	port     string = "3306"
	username string = "root"
	password string = "Bogia3110"
	database string = "chilindo"
)
var connectString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	username,
	password,
	host,
	port,
	database,
)

var DB *gorm.DB
var err error

func ConnectDatabase() {
	DB, err = gorm.Open(mysql.Open(connectString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err.Error())
	}
	if errConnect := DB.AutoMigrate(&entity.Administrator{}); errConnect != nil {
		panic(errConnect.Error())
	}
}

func GetDB() *gorm.DB {
	return DB
}

func init() {
	ConnectDatabase()
	log.Println("Connected to database...")
}

func CloseDatabase(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection")
	}
	dbSQL.Close()
}
