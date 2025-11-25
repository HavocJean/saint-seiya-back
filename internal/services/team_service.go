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
	var count int64
	if err := database.DB.
		Model(&models.Team{}).
		Where("user_id = ?", req.UserID).
		Count(&count).Error; err != nil {
		return nil, fmt.Errorf("failed to get teams: %w", err)
	}

	if count >= 5 {
		return nil, fmt.Errorf("user already has five teams")
	}

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
	var count int64
	if err := database.DB.
		Model(&models.TeamKnight{}).
		Where("team_id = ?", req.TeamID).
		Count(&count).Error; err != nil {
		return nil, fmt.Errorf("failed to get knights in team: %w", err)
	}

	if count >= 6 {
		return nil, fmt.Errorf("this team already has six knights")
	}

	var exists int64
	database.DB.
		Model(&models.TeamKnight{}).
		Where("team_id = ? AND knight_id = ?", req.TeamID, req.KnightID).
		Count(&exists)

	if exists > 0 {
		return nil, fmt.Errorf("this knight is already in the team")
	}

	teamKnight := &models.TeamKnight{
		TeamID:   req.TeamID,
		KnightID: req.KnightID,
	}

	if err := database.DB.Create(teamKnight).Error; err != nil {
		return nil, fmt.Errorf("failed to add knight in team: %w", err)
	}

	return teamKnight, nil
}
