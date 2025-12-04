package knight

import (
	"saint-seiya-back/internal/application/knight/dto"
	"saint-seiya-back/internal/domain/knight"
)

type GetKnightsUseCase struct {
	Repository knight.Repository
}

type GetKnightsInput struct {
	Page  int
	Limit int
	Rank  string
	Name  string
}

func NewGetKnightsUseCase(repo knight.Repository) *GetKnightsUseCase {
	return &GetKnightsUseCase{
		Repository: repo,
	}
}

func (u *GetKnightsUseCase) Execute(input GetKnightsInput) ([]dto.GetKnightsResponse, error) {
	knights, err := u.Repository.GetKnights(input.Page, input.Limit, input.Rank, input.Name)
	if err != nil {
		return nil, err
	}

	var result []dto.GetKnightsResponse
	for _, k := range knights {
		result = append(result, dto.GetKnightsResponse{
			ID:       k.ID,
			Name:     k.Name,
			Rank:     k.Rank,
			ImageURL: k.ImageURL,
		})
	}

	return result, nil
}
