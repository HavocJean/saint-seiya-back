package team

import (
	"errors"
	"saint-seiya-back/internal/application/team/dto"
	"saint-seiya-back/internal/domain/team"
)

type AddKnightToTeamUseCase struct {
	Repository team.Repository
}

func NewAddKnightToTeamUseCase(repo team.Repository) *AddKnightToTeamUseCase {
	return &AddKnightToTeamUseCase{
		Repository: repo,
	}
}

type AddKnightToTeamInput struct {
	TeamID   uint
	KnightID uint
}

func (ak *AddKnightToTeamUseCase) Execute(input AddKnightToTeamInput) (*dto.AddKnightToTeamResponse, error) {
	count, err := ak.Repository.CountKnightsByTeamID(input.TeamID)
	if err != nil {
		return nil, errors.New("failed to count knights in team")
	}

	if count >= 6 {
		return nil, errors.New("this team alreay has six knights")
	}

	exists, err := ak.Repository.KnightExistsInTeam(input.TeamID, input.KnightID)
	if err != nil {
		return nil, errors.New("failed to check if knight exists in team")
	}

	if exists {
		return nil, errors.New("this knight is already in the team")
	}

	teamKnightdomain := &team.TeamKnightDomain{
		TeamID:   input.TeamID,
		KnightID: input.KnightID,
	}

	createdTeamKnight, err := ak.Repository.AddKnightToTeam(teamKnightdomain)
	if err != nil {
		return nil, errors.New("failed to add knight to team")
	}

	return &dto.AddKnightToTeamResponse{
		ID:       createdTeamKnight.ID,
		TeamID:   createdTeamKnight.TeamID,
		KnightID: createdTeamKnight.KnightID,
	}, nil
}
