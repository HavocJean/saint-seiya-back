package controllers

import (
	"net/http"
	"saint-seiya-back/internal/application/team"
	"saint-seiya-back/internal/application/team/dto"
	"saint-seiya-back/internal/responses"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type TeamController struct {
	createTeamUseCase       *team.CreateTeamUseCase
	addKnightToTeamUseCase  *team.AddKnightToTeamUseCase
	deleteTeamUseCase       *team.DeleteTeamUseCase
	deleteTeamKnightUseCase *team.DeleteKnightToTeamUseCase
	getPublicTeamsUseCase   *team.GetPublicTeamsUseCase
}

func NewTeamController(
	createTeamUseCase *team.CreateTeamUseCase,
	addKnightToTeamUseCase *team.AddKnightToTeamUseCase,
	deleteTeamUseCase *team.DeleteTeamUseCase,
	deleteTeamKnightUseCase *team.DeleteKnightToTeamUseCase,
	getPublicTeamsUseCase *team.GetPublicTeamsUseCase,
) *TeamController {
	return &TeamController{
		createTeamUseCase:       createTeamUseCase,
		addKnightToTeamUseCase:  addKnightToTeamUseCase,
		deleteTeamUseCase:       deleteTeamUseCase,
		deleteTeamKnightUseCase: deleteTeamKnightUseCase,
		getPublicTeamsUseCase:   getPublicTeamsUseCase,
	}
}

func (tc *TeamController) CreateTeam(c *gin.Context) {
	var req dto.CreateTeamRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			responses.ValidationError(c, http.StatusBadRequest, err)
			return
		}
	}

	userID, exists := c.Get("user_id")
	if !exists {
		responses.Error(c, http.StatusBadRequest, "User ID not found in token", "Unauthorized")
		return
	}

	result, err := tc.createTeamUseCase.Execute(team.CreateTeamInput{
		Name:     req.Name,
		UserID:   userID.(uint),
		IsPublic: req.IsPublic,
	})

	if err != nil {
		responses.Error(c, http.StatusBadRequest, "Failed to create team", err.Error())
		return
	}

	responses.Success(c, http.StatusCreated, "Team created successfully", result)
}

func (tc *TeamController) AddKnightToTeam(c *gin.Context) {
	teamIDstring := c.Param("id")
	teamID, err := strconv.ParseUint(teamIDstring, 10, 64)

	if err != nil {
		responses.Error(c, http.StatusBadRequest, "ID invalid", "ID must be a number")
	}

	var req dto.AddKnightToTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			responses.ValidationError(c, http.StatusBadRequest, err)
			return
		}
		responses.Error(c, http.StatusBadRequest, "Invalid JSON sent", err.Error())
		return
	}

	result, err := tc.addKnightToTeamUseCase.Execute(team.AddKnightToTeamInput{
		TeamID:   uint(teamID),
		KnightID: req.KnightID,
	})

	if err != nil {
		responses.Error(c, http.StatusBadRequest, "Failed to add knight in team", err.Error())
		return
	}

	responses.Success(c, http.StatusCreated, "Add knight to the team successfuly", result)
}

func (tc *TeamController) DeleteTeam(c *gin.Context) {
	teamIDString := c.Param("id")
	teamID, err := strconv.ParseUint(teamIDString, 10, 64)
	if err != nil {
		responses.Error(c, http.StatusBadRequest, "ID invalid", "ID must be a number")
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		responses.Error(c, http.StatusUnauthorized, "User ID not found in token", "Unauthorized")
		return
	}

	err = tc.deleteTeamUseCase.Execute(team.DeleteTeamInput{
		TeamID: uint(teamID),
		UserID: userID.(uint),
	})

	if err != nil {
		responses.Error(c, http.StatusBadRequest, "Failed to delete team", err.Error())
		return
	}

	responses.Deleted(c, http.StatusNoContent, "Delete team successfully")
}

func (tc *TeamController) DeleteTeamKnight(c *gin.Context) {
	teamIDString := c.Param("teamId")
	teamID, errTeamID := strconv.ParseUint(teamIDString, 10, 64)

	knightIDString := c.Param("knightId")
	knightID, errKnightID := strconv.ParseUint(knightIDString, 10, 64)

	if errTeamID != nil || errKnightID != nil {
		responses.Error(c, http.StatusBadRequest, "ID invalid", "ID must be a number")
		return
	}

	err := tc.deleteTeamKnightUseCase.Execute(team.DeleteKnightToTeamInput{
		TeamID:   uint(teamID),
		KnightID: uint(knightID),
	})

	if err != nil {
		responses.Error(c, http.StatusBadRequest, "Failed to delete knight from team", err.Error())
		return
	}

	responses.Deleted(c, http.StatusNoContent, "Delete knight from team successfully")
}

func (tc *TeamController) GetPublicTeams(c *gin.Context) {
	page := 1
	limit := 20

	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}

	if l := c.Query("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 {
			limit = parsed
		}
	}

	result, err := tc.getPublicTeamsUseCase.Execute(team.GetPublicInput{
		Page:  page,
		Limit: limit,
	})

	if err != nil {
		responses.Error(c, http.StatusInternalServerError, "Error to get public teams", err.Error())
		return
	}

	responses.Success(c, http.StatusOK, "Public teams retrieved successfully", result)
}
