package controllers

import (
	"net/http"
	"strconv"

	"saint-seiya-awakening/internal/dto"
	"saint-seiya-awakening/internal/responses"
	"saint-seiya-awakening/internal/services"

	"github.com/gin-gonic/gin"
)

var knightService = services.NewKnightService()

func CreateKnight(c *gin.Context) {
	var req dto.CreateKnightRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.Error(c, http.StatusBadRequest, "Invalid JSON sent", err.Error())
		return
	}

	knight, err := knightService.CreateKnight(&req)
	if err != nil {
		responses.Error(c, http.StatusInternalServerError, "Failed to create knight", err.Error())
		return
	}

	responses.Success(c, http.StatusCreated, "Knight created successfully", knight)
}

func GetKnights(c *gin.Context) {
	page := 1
	limit := 20
	rank := c.Query("rank")

	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}

	if l := c.Query("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 && parsed <= 100 {
			limit = parsed
		}
	}

	knights, err := knightService.GetAllKnights(page, limit, rank)
	if err != nil {
		responses.Error(c, http.StatusInternalServerError, "Error internal to get knights", err.Error())
		return
	}

	responses.Success(c, http.StatusOK, "Knights retrieved successfully", knights)
}

func GetKnightById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		responses.Error(c, http.StatusBadRequest, "ID invalid", "ID must be a number")
		return
	}

	knight, err := knightService.GetKnightByID(uint(id))
	if err != nil {
		responses.Error(c, http.StatusNotFound, "Knight not found", err.Error())
		return
	}

	responses.Success(c, http.StatusOK, "Knight found", knight)
}

func CreateKnightSkill(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		responses.Error(c, http.StatusBadRequest, "ID invalid", "ID must be a number")
		return
	}

	var req dto.CreateKnightSkillRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.Error(c, http.StatusBadRequest, "Invalid JSON sent", err.Error())
		return
	}

	skill, err := knightService.CreateKnightSkill(uint(id), &req)
	if err != nil {
		responses.Error(c, http.StatusInternalServerError, "Failed to create knight skill", err.Error())
	}

	responses.Success(c, http.StatusCreated, "Knight skill created successfully", skill)
}
