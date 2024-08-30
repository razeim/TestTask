package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/razeim/testTask/internal/storage"
	"github.com/razeim/testTask/internal/token"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"

	}

	db, err := storage.DBSet()
	if err != nil {
		log.Fatal("", err)
	}
	defer db.Close()

	if err := storage.CreateUsersTable(db); err != nil {
		log.Fatal("Пользователи не создались:", err)
	}
	if err := storage.CreateTokensTable(db); err != nil {
		log.Fatal("Токены не создались:", err)
	}
	if err := storage.SeedUsersTable(db); err != nil {
		log.Fatal("Не получилось заполнить пользователей:", err)
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.POST("/generate-tokens", token.TokenGenerate)
	router.POST("/refresh-tokens", token.RefreshToken)

	log.Fatal(router.Run(":" + port))
}
