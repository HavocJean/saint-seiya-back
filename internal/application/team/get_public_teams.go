package team

import (
	"saint-seiya-back/internal/application/team/dto"
	"saint-seiya-back/internal/domain/team"
)

type GetPublicTeamsUseCase struct {
	Repository team.Repository
}

type GetPublicInput struct {
	Page  int
	Limit int
}

func NewGetPublicTeamsUseCase(repo team.Repository) *GetPublicTeamsUseCase {
	return &GetPublicTeamsUseCase{
		Repository: repo,
	}
}

func (gt *GetPublicTeamsUseCase) Execute(input GetPublicInput) ([]dto.GetPublicTeamsResponse, error) {
	teams, err := gt.Repository.GetPublicTeams(input.Page, input.Limit)

	if err != nil {
		return nil, err
	}

	result := make([]dto.GetPublicTeamsResponse, len(teams))
	for i, t := range teams {
		knights := make([]dto.TeamKnightResponse, len(t.Knights))
		for j, k := range t.Knights {
			imageURL := ""
			if k.ImageURL != nil {
				imageURL = *k.ImageURL
			}

			knights[j] = dto.TeamKnightResponse{
				KnightID: k.KnightID,
				Name:     k.Name,
				ImageURL: imageURL,
			}
		}

		result[i] = dto.GetPublicTeamsResponse{
			ID:      t.ID,
			Name:    t.Name,
			Knights: knights,
		}
	}

	return result, nil
}
