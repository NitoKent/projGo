package config

import (
	"fmt"
	"log"
	"os"

	"example.com/m/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDatabase() {
	// Загружаем переменные среды
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Настройка базы данных
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	defaultDB := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=disable", host, user, password, port)
	db, err := gorm.Open(postgres.Open(defaultDB), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to default database: %v", err)
	}
	sqlDB, _ := db.DB()

	// Проверяем, существует ли база данных
	var exists bool
	err = sqlDB.QueryRow(`SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = $1)`, dbname).Scan(&exists)
	if err != nil {
		log.Fatalf("Error checking if database exists: %v", err)
	}

	// Создаём базу данных, если её нет
	if !exists {
		_, err = sqlDB.Exec("CREATE DATABASE " + dbname)
		if err != nil {
			log.Fatalf("Error creating database: %v", err)
		}
		fmt.Printf("Database '%s' created successfully\n", dbname)
	} else {
		fmt.Printf("Database '%s' already exists\n", dbname)
	}

	sqlDB.Close()

	// Подключаемся к новой базе данных
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to '%s' database: %v", dbname, err)
	}
	fmt.Printf("Successfully connected to '%s' database\n", dbname)

	// Миграция таблицы для модели Book
	err = DB.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	fmt.Println("Database migration completed successfully")
}
