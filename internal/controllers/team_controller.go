package controllers

import (
	"net/http"
	"saint-seiya-awakening/internal/dto"
	"saint-seiya-awakening/internal/responses"
	"saint-seiya-awakening/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

var teamService = services.NewTeamService()

func CreateTeam(c *gin.Context) {
	var req dto.CreateTeamRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.Error(c, http.StatusBadRequest, "Invalid JSON sent", err.Error())
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		responses.Error(c, http.StatusUnauthorized, "User ID not found in token", "Unauthorized")
		return
	}

	req.UserID = userID.(uint)

	team, err := teamService.CreateTeam(&req)
	if err != nil {
		responses.Error(c, http.StatusBadRequest, "Failed to create team", err.Error())
		return
	}

	responses.Success(c, http.StatusCreated, "Team created successfully", team)
}

func AddKnightToTeam(c *gin.Context) {
	idString := c.Param("id")
	teamId, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		responses.Error(c, http.StatusBadRequest, "ID invalid", "ID must be a number")
		return
	}

	var req dto.CreateTeamKnight
	req.TeamID = uint(teamId)

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.Error(c, http.StatusBadRequest, "Invalid JSON sent", err.Error())
		return
	}

	team, err := teamService.AddKnightToTeam(&req)
	if err != nil {
		responses.Error(c, http.StatusBadRequest, "Failed to add knight in team", err.Error())
		return
	}

	responses.Success(c, http.StatusCreated, "Add knight to the team successfully", team)
}

func DeleteTeam(c *gin.Context) {
	idString := c.Param("id")
	teamId, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		responses.Error(c, http.StatusBadRequest, "ID invalid", "ID must be a number")
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		responses.Error(c, http.StatusUnauthorized, "User ID not found in token", "Unauthorized")
		return
	}

	team, err := teamService.DeleteTeam(uint(teamId), userID.(uint))
	if err != nil {
		responses.Error(c, http.StatusBadRequest, "Failed to delete knight in team", err.Error())
		return
	}

	responses.Success(c, http.StatusNoContent, "Delete team successfully", team)
}

func DeleteTeamKnight(c *gin.Context) {
	teamIdStr := c.Param("teamId")
	teamId, errTeamId := strconv.ParseUint(teamIdStr, 10, 64)

	knightIdStr := c.Param("knightId")
	knightId, errKnightId := strconv.ParseUint(knightIdStr, 10, 64)

	if errTeamId != nil || errKnightId != nil {
		responses.Error(c, http.StatusBadRequest, "ID invalid", "ID must be a number")
		return
	}

	team, err := teamService.DeleteTeamKnight(uint(teamId), uint(knightId))
	if err != nil {
		responses.Error(c, http.StatusBadRequest, "ID invalid", "ID must be a number")
		return
	}

	responses.Success(c, http.StatusNoContent, "Delete knight from team successfully", team)
}
