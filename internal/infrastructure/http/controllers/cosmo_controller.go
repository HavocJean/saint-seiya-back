package controllers

import (
	"net/http"
	"saint-seiya-back/internal/application/cosmo"
	"saint-seiya-back/internal/responses"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CosmoController struct {
	getCosmosUseCase    *cosmo.GetCosmosUseCase
	getCosmoByIdUseCase *cosmo.GetCosmoByIdUseCase
}

func NewCosmoController(
	getCosmosUseCase *cosmo.GetCosmosUseCase,
	getCosmoByIdUseCase *cosmo.GetCosmoByIdUseCase,
) *CosmoController {
	return &CosmoController{
		getCosmosUseCase:    getCosmosUseCase,
		getCosmoByIdUseCase: getCosmoByIdUseCase,
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
