package controllers

import (
	"net/http"
	"saint-seiya-awakening/internal/database"
	"saint-seiya-awakening/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateCosmoRequest struct {
	Name           string                  `json:"name" binding:"required"`
	Rank           string                  `json:"rank" binding:"required"`
	SetBonus       string                  `json:"set_bonus" binding:"required"`
	ImageURL       *string                 `json:"image_url"`
	BaseAttributes []models.CosmoAttribute `json:"base_attributes" binding:"required"`
}

func CreateCosmo(c *gin.Context) {
	var req CreateCosmoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cosmo := models.Cosmo{
		Name:           req.Name,
		Rank:           req.Rank,
		SetBonus:       req.SetBonus,
		ImageURL:       req.ImageURL,
		BaseAttributes: req.BaseAttributes,
	}

	if err := database.DB.Create(&cosmo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cosmo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Cosmo created successfully",
	})
}
