package storage

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestDBSet(t *testing.T) {
	_, err := DBSet()
	if err != nil {
		t.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
}

func TestCreateUsersTable(t *testing.T) {
	db, err := DBSet()
	if err != nil {
		t.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer db.Close()

	err = CreateUsersTable(db)
	if err != nil {
		t.Fatalf("Не удалось создать таблицу : %v", err)
	}
}

func TestCreateTokensTable(t *testing.T) {
	db, err := DBSet()
	if err != nil {
		t.Fatalf("Не удвлось подключиться к базе данных: %v", err)
	}
	defer db.Close()

	err = CreateTokensTable(db)
	if err != nil {
		t.Fatalf(":Не удалось создать таблицу %v", err)
	}
}

func TestSeedUsersTable(t *testing.T) {
	db, err := DBSet()
	if err != nil {
		t.Fatalf("Не удвлось подключиться к базе данных: %v", err)
	}
	defer db.Close()

	err = SeedUsersTable(db)
	if err != nil {
		t.Fatalf("Не удвлось заполнить таблицу: %v", err)
	}
}
