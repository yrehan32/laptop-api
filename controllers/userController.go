package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gitlab.com/yrehan32/laptop-api/initializers"
	"gitlab.com/yrehan32/laptop-api/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"message": "Invalid request body",
			"data": nil,
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"message": "Failed to hash password",
			"data": nil,
		})

		return
	}

	user := models.User{
		Name: body.Name,
		Email: body.Email,
		Password: string(hash),
	}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"message": "Failed to create user",
			"data": nil,
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"error": false,
		"message": "User created successfully",
		"data": user,
	})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"message": "Invalid request body",
			"data": nil,
		})

		return
	}

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"message": "Invalid email",
			"data": nil,
		})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"message": "Invalid password",
			"data": nil,
		})

		return
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("APP_KEY")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"message": "Failed to create token",
			"data": nil,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"message": "Login successful",
		"data": gin.H{
			"token": tokenString,
		},
	})
}

func GetUser(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"message": "User fetched successfully",
		"data": user,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name     string `json:"name" binding:"required"`
	}

	binding := c.Bind(&body)
	if binding != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"message": "Invalid request body",
			"data": nil,
		})

		return
	}

	var user models.User
	initializers.DB.First(&user, id)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": true,
			"message": "User with given ID not found",
			"data": nil,
		})

		return
	}

	initializers.DB.Model(&user).Updates(models.User{
		Name: body.Name,
	})

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"message": "User updated successfully",
		"data": nil,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	initializers.DB.First(&user, id)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": true,
			"message": "User with given ID not found",
			"data": nil,
		})

		return
	}

	initializers.DB.Delete(&models.User{}, id)

	if initializers.DB.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"message": "Failed to delete user",
			"data": nil,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"message": "User deleted successfully",
		"data": nil,
	})
}