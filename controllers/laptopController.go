package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/yrehan32/laptop-api/initializers"
	"gitlab.com/yrehan32/laptop-api/models"
	"gorm.io/gorm"
)

func GetAllLaptop(c *gin.Context) {
	var laptops []models.Laptop

	result := initializers.DB.Find(&laptops)
	
	if result.RowsAffected == 0 {
		ErrorResponse(c, http.StatusNotFound, "No laptop found")
		return
	}

	if result.Error != nil {
		ErrorResponse(c, http.StatusInternalServerError, "Failed to get laptops")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"message": "Laptop data fetched successfully.",
		"data": laptops,
	})
}

func GetLaptopById(c *gin.Context) {
	id := c.Param("id")
	var laptop models.Laptop

	result := initializers.DB.First(&laptop, id)
	
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ErrorResponse(c, http.StatusNotFound, "Laptop not found")
		} else {
			ErrorResponse(c, http.StatusInternalServerError, "Failed to get laptop")
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"message": "Laptop fetched successfully",
		"data": laptop,
	})
}

func CreateLaptop(c *gin.Context) {
	var body struct {
		Name        string `json:"name" binding:"required"`
		Brand       string `json:"brand" binding:"required"`
		Processor   string `json:"processor" binding:"required"`
		RAM         int    `json:"ram" binding:"required"`
		Storage     int    `json:"storage" binding:"required"`
		Display     string `json:"display" binding:"required"`
		Graphics    string `json:"graphics" binding:"required"`
		OS          string `json:"os"`
		Battery     string `json:"battery"`
		Keyboard    string `json:"keyboard"`
		Description string `json:"description"`
		Price       int    `json:"price" binding:"required"`
	}

	if c.Bind(&body) != nil {
		ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	laptop := models.Laptop{
		Name:        body.Name,
		Brand:       body.Brand,
		Processor:   body.Processor,
		RAM:         body.RAM,
		Storage:     body.Storage,
		Display:     body.Display,
		Graphics:    body.Graphics,
		OS:          body.OS,
		Battery:     body.Battery,
		Keyboard:    body.Keyboard,
		Description: body.Description,
		Price:       body.Price,
	}

	result := initializers.DB.Create(&laptop)

	if result.Error != nil {
		ErrorResponse(c, http.StatusInternalServerError, "Failed to create laptop")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"error":   false,
		"message": "Laptop created successfully",
		"data":    laptop,
	})
}

func UpdateLaptop(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name        string `json:"name" binding:"required"`
		Brand       string `json:"brand" binding:"required"`
		Processor   string `json:"processor" binding:"required"`
		RAM         int    `json:"ram" binding:"required"`
		Storage     int    `json:"storage" binding:"required"`
		Display     string `json:"display" binding:"required"`
		Graphics    string `json:"graphics" binding:"required"`
		OS          string `json:"os"`
		Battery     string `json:"battery"`
		Keyboard    string `json:"keyboard"`
		Description string `json:"description"`
		Price       int    `json:"price" binding:"required"`
	}

	binding := c.Bind(&body)
	if binding != nil {
		ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}
	
	var laptop models.Laptop
	initializers.DB.First(&laptop, id)

	if laptop.ID == 0 {
		ErrorResponse(c, http.StatusNotFound, "Laptop not found")
		return
	}

	initializers.DB.Model(&laptop).Updates(models.Laptop{
		Name:        body.Name,
		Brand:       body.Brand,
		Processor:   body.Processor,
		RAM:         body.RAM,
		Storage:     body.Storage,
		Display:     body.Display,
		Graphics:    body.Graphics,
		OS:          body.OS,
		Battery:     body.Battery,
		Keyboard:    body.Keyboard,
		Description: body.Description,
		Price:       body.Price,
	})

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Laptop updated successfully",
		"data":    laptop,
	})
}

func DeleteLaptop(c *gin.Context)  {
	id := c.Param("id")
	
	var laptop models.Laptop
	initializers.DB.First(&laptop, id)

	if laptop.ID == 0 {
		ErrorResponse(c, http.StatusNotFound, "Laptop not found")
		return
	}

	initializers.DB.Delete(&laptop, id)

	if initializers.DB.Error != nil {
		ErrorResponse(c, http.StatusInternalServerError, "Failed to delete laptop")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"message": "Laptop deleted successfully",
		"data": nil,
	})
}
