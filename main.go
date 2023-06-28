package main

import (
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
	
	r.Run() // listen and serve on 0.0.0.0:8080
}