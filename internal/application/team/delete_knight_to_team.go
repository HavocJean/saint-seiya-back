package team

import (
	"errors"
	"saint-seiya-back/internal/domain/team"
)

type DeleteKnightToTeamUseCase struct {
	Repository team.Repository
}

func NewDeleteKnightToTeamUseCase(repo team.Repository) *DeleteKnightToTeamUseCase {
	return &DeleteKnightToTeamUseCase{
		Repository: repo,
	}
}

type DeleteKnightToTeamInput struct {
	TeamID   uint
	KnightID uint
}

func (dtk *DeleteKnightToTeamUseCase) Execute(input DeleteKnightToTeamInput) error {
	err := dtk.Repository.DeleteKnightToTeam(input.TeamID, input.KnightID)
	if err != nil {
		return errors.New("team knight not found")
	}

	return nil
}
