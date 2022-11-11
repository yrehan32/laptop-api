package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gitlab.com/yrehan32/laptop-api/initializers"
	"gitlab.com/yrehan32/laptop-api/models"
)

func RequireAuth(c *gin.Context) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": true,
			"message": "Missing Authorization header",
			"data": nil,
		})
	}

	tokenString := authHeader[len(BEARER_SCHEMA):]

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
	
		return []byte(os.Getenv("APP_KEY")), nil
	})
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": true,
				"message": "Expired token",
				"data": nil,
			})
		}

		var user models.User
		initializers.DB.First(&user, claims["sub"].(float64))

		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": true,
				"message": "Invalid token",
				"data": nil,
			})
		}

		// If valid, set user to context
		c.Set("user", user)
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": true,
			"message": "Unauthorized",
			"data": nil,
		})
	}
}