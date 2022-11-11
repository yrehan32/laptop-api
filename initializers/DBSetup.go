package initializers

import "gitlab.com/yrehan32/laptop-api/models"

func AutoMigrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Laptop{})
}