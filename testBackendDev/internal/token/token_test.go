package token

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/generate-tokens", TokenGenerate)
	router.POST("/refresh-tokens", RefreshToken)
	return router
}

func TestTokenGenerate(t *testing.T) {
	router := setupRouter()
	req := httptest.NewRequest("POST", "/generate-tokens?id=1e4b7b1f8-f4d4-4e1a-a7cb-65e0a2d59c92", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
func TestRefreshToken(t *testing.T) {
	router := setupRouter()
	req := httptest.NewRequest("POST", "/refresh-tokens?id=e4b7b1f8-f4d4-4e1a-a7cb-65e0a2d59c92", nil)
	req.PostForm = map[string][]string{
		"refresh_token": {""},
	}

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
