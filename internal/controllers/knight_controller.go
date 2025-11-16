package controllers

import (
	"net/http"
	"saint-seiya-awakening/internal/database"
	"saint-seiya-awakening/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateKnightSkillRequest struct {
	Name        string                    `json:"name" binding:"required"`
	Type        string                    `json:"type" binding:"required"`
	ImageURL    string                    `json:"image_url"`
	Description string                    `json:"description" binding:"required"`
	Levels      []models.KnightSkillLevel `json:"levels" binding:"required"`
}

func CreateKnight(c *gin.Context) {
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
		"message": "Knight created successfully",
	})
}

func GetKnights(c *gin.Context) {
	var knights []models.Knight
	if err := database.DB.Find(&knights).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve knights"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"knights": knights,
	})
}

func GetKnightById(c *gin.Context) {
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

func CreateKnightSkill(c *gin.Context) {
	var req CreateKnightSkillRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idStr := c.Param("id")
	idUint64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid knight id"})
		return
	}
	knightID := uint(idUint64)

	skill := models.KnightSkill{
		KnightID:    knightID,
		Name:        req.Name,
		Type:        req.Type,
		ImageURL:    req.ImageURL,
		Description: req.Description,
		Levels:      req.Levels,
	}

	if err := database.DB.Create(&skill).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create knight skill"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Knight skill created successfully",
	})
}
