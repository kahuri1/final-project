package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kahuri1/final-project/pkg/model"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var jwtKey = []byte(viper.GetString("JWT_SECRET_KEY"))

func (h *Handler) Auth(c *gin.Context) {
	var creds *model.Auth
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный запрос"})
		return
	}
	expectedPassword := viper.GetString("TODO_PASSWORD")
	if creds.Password != expectedPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный пароль"})
		return
	}
	expirationTime := time.Now().Add(8 * time.Hour)
	claims := &model.Claims{
		PasswordHash: creds.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать токен"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
