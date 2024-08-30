package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Application struct {
	db *sql.DB
}

func DBSet() (*sql.DB, error) {
	connString := "postgres://postgres:parol@db:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Успешное подключение к базе данных")

	return db, nil
}

func CreateUsersTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		email VARCHAR(255) UNIQUE NOT NULL,
		created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	log.Println("Таблица пользователей создана или уже существует")
	return nil
}

func CreateTokensTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS tokens (
		id SERIAL PRIMARY KEY,
		user_id UUID NOT NULL,
		refresh_token TEXT NOT NULL,
		user_ip TEXT NOT NULL,
		created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
		expires_at TIMESTAMPTZ NOT NULL
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	log.Println("Таблица токенов создана или уже существует")
	return nil
}

func SeedUsersTable(db *sql.DB) error {
	query := `
	INSERT INTO users (email) VALUES
		('1@gmail.com'),
		('2@gmail.com'),
		('3@gmail.com')
	ON CONFLICT (email) DO NOTHING;
	`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
