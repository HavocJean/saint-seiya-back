package controllers

import (
	"net/http"
	"saint-seiya-awakening/internal/database"
	"saint-seiya-awakening/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateCosmoRequest struct {
	Name              string                  `json:"name" binding:"required"`
	Rank              string                  `json:"rank" binding:"required"`
	Color             models.CosmoColor       `json:"color" binding:"required"`
	SetBonusValue     float64                 `json:"set_bonus" binding:"required"`
	SetBonusName      string                  `json:"set_bonus_name" binding:"required"`
	SetBonusIsPercent bool                    `json:"set_bonus_is_percent" binding:"required"`
	ImageURL          *string                 `json:"image_url"`
	BaseAttributes    []models.CosmoAttribute `json:"base_attributes" binding:"required"`
}

func CreateCosmo(c *gin.Context) {
	var req CreateCosmoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cosmo := models.Cosmo{
		Name:              req.Name,
		Rank:              req.Rank,
		Color:             req.Color,
		SetBonusValue:     req.SetBonusValue,
		SetBonusName:      req.SetBonusName,
		SetBonusIsPercent: req.SetBonusIsPercent,
		ImageURL:          req.ImageURL,
		BaseAttributes:    req.BaseAttributes,
	}

	if err := database.DB.Create(&cosmo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cosmo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Cosmo created successfully",
	})
}

func GetCosmos(c *gin.Context) {
	var cosmos []models.Cosmo
	if err := database.DB.Find(&cosmos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cosmos"})
	}

	c.JSON(http.StatusOK, gin.H{
		"cosmos": cosmos,
	})
}

func GetCosmosById(c *gin.Context) {
	id := c.Param("id")
	var cosmo models.Cosmo
	if err := database.DB.First(&cosmo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cosmo not found"})
	}

	c.JSON(http.StatusOK, gin.H{
		"cosmo": cosmo,
	})
}
