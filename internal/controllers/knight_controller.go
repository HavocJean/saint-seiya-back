package controller

import (
	"net/http"
	"saint-seiya-awakening/internal/database"
	"saint-seiya-awakening/internal/models"

	"github.com/gin-gonic/gin"
)

func createKnight(c *gin.Context) {
	var knight models.Knight
	if err := c.ShouldBindJSON(&knight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&knight).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create knight"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Knight created successfully"
	})
}

func getKnights(c *gin.Context) {
	var knights []models.Knight
	if err := database.DB.Find(&knights).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve knights"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"knights": knights,
	})
}

func getKnightById(c *gin.Context) {
	id := c.Param("id")
	var knight models.Knight
	if err := database.DB.First(&knight, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Knight not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"knight": knight,
	})
}