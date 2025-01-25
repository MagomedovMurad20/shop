package main

import (
	"log"
	"shop/internal/db"
)

func main() {
	db.InitDB()
	log.Println("Приложение запущено")

}
