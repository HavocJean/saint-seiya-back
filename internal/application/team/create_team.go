package team

import (
	"errors"
	"saint-seiya-back/internal/application/team/dto"
	"saint-seiya-back/internal/domain/team"
)

type CreateTeamUseCase struct {
	Repository team.Repository
}

func NewCreateTeamUseCase(repo team.Repository) *CreateTeamUseCase {
	return &CreateTeamUseCase{
		Repository: repo,
	}
}

type CreateTeamInput struct {
	Name     string
	UserID   uint
	IsPublic bool
}

func (ct *CreateTeamUseCase) Execute(input CreateTeamInput) (*dto.CreateTeamResponse, error) {
	count, err := ct.Repository.CountByUserID(input.UserID)
	if err != nil {
		return nil, errors.New("failed to count user teams")
	}

	if count >= 5 {
		return nil, errors.New("user already has five teams")
	}

	teamDomain := &team.TeamDomain{
		Name:     input.Name,
		UserID:   input.UserID,
		IsPublic: input.IsPublic,
	}

	createdTeam, err := ct.Repository.Create(teamDomain)
	if err != nil {
		return nil, errors.New("failed to create team")
	}

	return &dto.CreateTeamResponse{
		ID:       createdTeam.ID,
		Name:     createdTeam.Name,
		UserID:   createdTeam.ID,
		IsPublic: createdTeam.IsPublic,
	}, nil
}
