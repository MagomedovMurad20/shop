package db

import (
	"database/sql"
	"fmt"
	"log"
	"shop/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	cfg := config.LoadConfig()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	log.Printf("Подключение к базе данных: host=%s, port=%s, user=%s, dbname=%s", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Не удалось проверить подключение к базе данных: %v", err)
	}

	log.Println("Успешное подключение к базе данных!")
}
