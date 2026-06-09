package main

import (
	"context"
	"fmt"
	"git/db"
	sql "git/testSQL"
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

	if err := sql.CreateUsersTable(ctx, pool); err != nil {
		log.Fatalf("Ошибка создания таблицы users: %v", err)
	}

	fmt.Println("Таблица users готова")

	err = sql.CreateUsersTable(ctx, pool)
	if err != nil {
		log.Fatalf("Ошибка создания таблицы users: %v", err)
	}

	usersByID, err := sql.GetUsersByID(ctx, pool)
	if err != nil {
		log.Fatalf("Ошибка получения пользователей: %v", err)
	}

	fmt.Println(usersByID)

	users, err := sql.GetUsers(ctx, pool)

	fmt.Println(users[0])

}
