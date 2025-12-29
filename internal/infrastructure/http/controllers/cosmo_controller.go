package controllers

import (
	"net/http"
	"saint-seiya-back/internal/application/cosmo"
	"saint-seiya-back/internal/application/cosmo/dto"
	"saint-seiya-back/internal/responses"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CosmoController struct {
	getCosmosUseCase    *cosmo.GetCosmosUseCase
	getCosmoByIdUseCase *cosmo.GetCosmoByIdUseCase
	createCosmoUseCase  *cosmo.CreateCosmoUseCase
}

func NewCosmoController(
	getCosmosUseCase *cosmo.GetCosmosUseCase,
	getCosmoByIdUseCase *cosmo.GetCosmoByIdUseCase,
	createCosmoUseCase *cosmo.CreateCosmoUseCase,
) *CosmoController {
	return &CosmoController{
		getCosmosUseCase:    getCosmosUseCase,
		getCosmoByIdUseCase: getCosmoByIdUseCase,
		createCosmoUseCase:  createCosmoUseCase,
	}
}

func (cc *CosmoController) GetCosmos(c *gin.Context) {
	color := c.Query("color")

	result, err := cc.getCosmosUseCase.Execute(color)
	if err != nil {
		responses.Error(c, http.StatusInternalServerError, "Failed to retrieve cosmos", err.Error())
		return
	}

	responses.Success(c, http.StatusInternalServerError, "Cosmos retrieved successfully", result)
}

func (cc *CosmoController) GetCosmoById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		responses.Error(c, http.StatusBadRequest, "ID invalid", "ID must be a number")
		return
	}

	result, err := cc.getCosmoByIdUseCase.Execute(uint(id))
	if err != nil {
		responses.Error(c, http.StatusNotFound, "Cosmo not found", err.Error())
		return
	}

	responses.Success(c, http.StatusOK, "Cosmo found", result)
}

func (cc *CosmoController) CreateCosmo(c *gin.Context) {
	var req dto.CreateCosmoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			responses.ValidationError(c, http.StatusBadRequest, err)
			return
		}
		responses.Error(c, http.StatusBadRequest, "Invalid JSON sent", err.Error())
		return
	}

	result, err := cc.createCosmoUseCase.Execute(cosmo.CreateCosmoInput{
		Name:              req.Name,
		Rank:              req.Rank,
		Color:             req.Color,
		SetBonusValue:     req.SetBonusValue,
		SetBonusName:      req.SetBonusName,
		SetBonusIsPercent: req.SetBonusIsPercent,
		ImageURL:          req.ImageURL,
		BaseAttributes:    req.BaseAttributes,
	})

	if err != nil {
		responses.Error(c, http.StatusInternalServerError, "Failed to create cosmo", err.Error())
		return
	}

	responses.Success(c, http.StatusCreated, "Cosmo created successfully", result)
}
