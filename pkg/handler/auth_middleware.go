package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kahuri1/final-project/pkg/model"
	"github.com/spf13/viper"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Аутентификация требуется"})
			c.Abort()
			return
		}
		claims := &model.Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
			c.Abort()
			return
		}

		expectedPassword := viper.GetString("TODO_PASSWORD")
		if claims.PasswordHash != expectedPassword {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неправильный токен"})
			c.Abort()
			return
		}

		c.Next()
	}
}
