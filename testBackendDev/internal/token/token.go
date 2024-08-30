package token

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/razeim/testTask/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	UserID string `json: "user_id"`
	UserIP string `json: "user_ip"`
	jwt.RegisteredClaims
}

var secret = []byte("I_LOVE_MEDODS")

func TokenGenerate(c *gin.Context) {
	userIdentificator := c.Query("id")
	if userIdentificator == "" {
		c.JSON(http.StatusBadRequest, "Пустой Id")
		return
	}

	db, err := storage.DBSet()
	if err != nil {
		log.Fatal("", err)
	}
	defer db.Close()

	var userEmail string

	err = db.QueryRow("SELECT email FROM users WHERE id = $1", userIdentificator).Scan(&userEmail)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	exTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		UserID: userIdentificator,
		UserIP: c.ClientIP(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokString, err := token.SignedString(secret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не получилось сгенерировать токен"})
		return
	}

	refreshToken := base64.StdEncoding.EncodeToString(([]byte(fmt.Sprintf("%s:%d", userIdentificator, time.Now().Unix()))))
	hashRefreshToken, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.MinCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не получилось сгенерировать рефреш токен"})
		return
	}

	_, err = db.Exec("INSERT INTO tokens (user_id, refresh_token, user_ip, created_at, expires_at) VALUES ($1, $2, $3, $4, $5)",
		userIdentificator, hashRefreshToken, c.ClientIP(), time.Now(), exTime.Add(24*time.Hour))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось записать токен"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  tokString,
		"refresh_token": refreshToken,
	})
}

func RefreshToken(c *gin.Context) {
	userIdentificator := c.Query("id")
	if userIdentificator == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id is empty"})
		return
	}

	refreshToken := c.PostForm("refresh_token")
	if refreshToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "пустой токен"})
		return
	}

	db, err := storage.DBSet()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось подключиться к базе данных"})
		return
	}
	defer db.Close()

	var hashRefreshToken, userIP string

	err = db.QueryRow("SELECT refresh_token, user_ip FROM tokens WHERE user_id=$1 ORDER BY created_at DESC LIMIT 1", userIdentificator).Scan(&hashRefreshToken, &userIP)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен не найден"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashRefreshToken), []byte(refreshToken)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный рефреш токен"})
		return
	}

	if userIP != c.ClientIP() {
		// TODO: отправка email
		log.Printf("IP адрес изменился для пользователя %s. Старый IP: %s, Новый IP: %s\n", userIdentificator, userIP, c.ClientIP())
	}

	exTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		UserID: userIdentificator,
		UserIP: c.ClientIP(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokString, err := token.SignedString(secret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сгенерировать токен"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": tokString})
}
