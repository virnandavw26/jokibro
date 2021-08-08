package database

import (
	"fmt"
	"jokibro/app/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlConfig struct {
	DBName     string
	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string
}

func connectionMap() mysqlConfig {
	config := mysqlConfig{
		DBName:     os.Getenv("DB_NAME"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}

	return config
}

func assembleConfig() string {
	config := connectionMap().DBUsername + ":" +
		connectionMap().DBPassword + "@(" +
		connectionMap().DBHost + ":" +
		connectionMap().DBPort + ")/" +
		connectionMap().DBName + "?" +
		"parseTime=true"
	return config
}

func InitDB() *gorm.DB {
	var err error
	db, err := gorm.Open(mysql.Open(assembleConfig()), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Failed Connection Database")
	}

	return db
}

func MigrateTables(db *gorm.DB) {
	db.AutoMigrate(&models.Category{})
}
