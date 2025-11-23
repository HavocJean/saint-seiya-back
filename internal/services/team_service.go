package services

import (
	"fmt"
	"saint-seiya-awakening/internal/database"
	"saint-seiya-awakening/internal/dto"
	"saint-seiya-awakening/internal/models"
)

type TeamService struct{}

func NewTeamService() *TeamService {
	return &TeamService{}
}

func (s *TeamService) CreateTeam(req *dto.CreateTeamRequest) (*models.Team, error) {
	team := &models.Team{
		Name:     req.Name,
		UserID:   req.UserID,
		IsPublic: req.IsPublic,
	}

	if err := database.DB.Create(team).Error; err != nil {
		return nil, fmt.Errorf("failed to create team: %w", err)
	}

	return team, nil
}

func (s *TeamService) AddKnightToTeam(req *dto.CreateTeamKnight) (*models.TeamKnight, error) {
	teamKnight := &models.TeamKnight{
		TeamID:   req.TeamID,
		KnightID: req.KnightID,
	}

	if err := database.DB.Create(teamKnight).Error; err != nil {
		return nil, fmt.Errorf("failed to add knight in team: %w", err)
	}

	return teamKnight, nil
}
