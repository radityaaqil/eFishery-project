package config

import (
	"efishery-project/entities"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

type ConfigDB struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
	TimeZone string
}

func ConfigurationDB() *ConfigDB {
	//Load env
	err := godotenv.Load(".env")

	if err != nil {
		log.Panic("Failed to load env")
	}

	config := &ConfigDB{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		TimeZone: os.Getenv("DB_TIMEZONE"),
	}

	return config
}

func NewConnection(configdb *ConfigDB) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", configdb.Host, configdb.User, configdb.Password, configdb.DBName, configdb.Port, configdb.SSLMode, configdb.TimeZone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Failed to establish connection to database")
		return db, err
	}

	db.AutoMigrate(&entities.User{}, &entities.Product{}, &entities.ProductImages{}, &entities.Cart{}, &entities.Transaction{}, &entities.Transaction_Detail{})

	return db, nil
}
