package sql

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

func CreateUsersTable(ctx context.Context, pool *pgxpool.Pool) error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);
	`

	_, err := pool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("create users table: %w", err)
	}

	return nil
}

func GetUsers(ctx context.Context, pool *pgxpool.Pool) ([]User, error) {

	query := `
		SELECT id, name, email, created_at
		FROM users
		ORDER BY id;
	`

	rows, err := pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Получение из таблицы users - ошибка %w ", err)
	}
	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		var user User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("Ошибка скан пользователей %w", err)
		}

		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return users, nil

}

func GetUsersByID(ctx context.Context, pool *pgxpool.Pool) (map[int]User, error) {
	users, err := GetUsers(ctx, pool)
	if err != nil {
		return nil, err
	}

	usersByID := make(map[int]User)

	for _, user := range users {
		usersByID[user.ID] = user
	}

	return usersByID, nil
}
