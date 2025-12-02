package knight

import "saint-seiya-back/internal/application/knight/dto"

type GetKnightsUseCase struct {
	Repository Repository
}

type GetKnightsInput struct {
	Page  int
	Limit int
	Rank  string
	Name  string
}

func NewGetKnightUseCase(repo Repository) *GetKnightsUseCase {
	return &GetKnightsUseCase{
		Repository: repo,
	}
}

func (u *GetKnightsUseCase) Execute(input GetKnightsInput) ([]dto.GetKnightsResponse, error) {
	knights, err := u.Repository.FindAll(input.Page, input.Limit, input.Rank, input.Name)
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

	return results, nil
}
