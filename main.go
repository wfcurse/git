package main

import (
	"context"
	"fmt"
	"git/db"
	"log"
)

func main() {
	ctx := context.Background()

	cfg := db.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "123",
		Name:     "go_project_db",
		SSLMode:  "disable",
	}

	pool, err := db.Connect(ctx, cfg)
	if err != nil {
		log.Fatalf("Ошибка подключения %v", err)
	}
	defer pool.Close()
	fmt.Println("Успешное подключение")
}
