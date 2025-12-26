package team

import (
	"errors"
	"saint-seiya-back/internal/domain/team"
)

type DeleteTeamUseCase struct {
	Repository team.Repository
}

func NewDeleteTeamUseCase(repo team.Repository) *DeleteTeamUseCase {
	return &DeleteTeamUseCase{
		Repository: repo,
	}
}

type DeleteTeamInput struct {
	TeamID uint
	UserID uint
}

func (dt *DeleteTeamUseCase) Execute(input DeleteTeamInput) error {
	team, err := dt.Repository.GetByID(input.TeamID)
	if err != nil {
		return errors.New("team not found")
	}

	if team.UserID != input.UserID {
		return errors.New("unauthorized: team does not belong to user")
	}

	err = dt.Repository.Delete(input.TeamID, input.UserID)
	if err != nil {
		return errors.New("failed to delete team")
	}

	return nil
}
