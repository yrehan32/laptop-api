package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/yrehan32/laptop-api/controllers"
	"gitlab.com/yrehan32/laptop-api/initializers"
	"gitlab.com/yrehan32/laptop-api/middleware"
)

func init() {
	initializers.LoadEnv()
	initializers.DBConnect()
	initializers.AutoMigrate()
}

func main() {
	r := gin.Default()

	// validate the domain
	r.Use(ValidateDomainMiddleware(os.Getenv("HOST")))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"error": false,
			"message": "Welcome to Laptop API",
			"data": nil,
		})
	})
	
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	r.GET("/user", middleware.RequireAuth, controllers.GetUser)
	r.PUT("/user/:id", middleware.RequireAuth, controllers.UpdateUser)
	r.DELETE("/user/:id", middleware.RequireAuth, controllers.DeleteUser)

	r.GET("/laptop", middleware.RequireAuth, controllers.GetAllLaptop)
	r.GET("/laptop/:id", middleware.RequireAuth, controllers.GetLaptopById)
	r.POST("/laptop", middleware.RequireAuth, controllers.CreateLaptop)
	r.PUT("/laptop/:id", middleware.RequireAuth, controllers.UpdateLaptop)
	r.DELETE("/laptop/:id", middleware.RequireAuth, controllers.DeleteLaptop)
	
	r.Run(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
}

// validates the incoming request's domain
func ValidateDomainMiddleware(allowedDomain string) gin.HandlerFunc {
	return func(c *gin.Context) {
		host := c.Request.Host
		if !strings.Contains(host, allowedDomain) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": true,
				"message": "Please use the correct domain",
				"data": nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}