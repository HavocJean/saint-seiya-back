package cosmo

import (
	"saint-seiya-back/internal/application/cosmo/dto"
	"saint-seiya-back/internal/domain/cosmo"
)

type GetCosmosUseCase struct {
	Repository cosmo.Repository
}

func NewGetCosmosUseCase(repo cosmo.Repository) *GetCosmosUseCase {
	return &GetCosmosUseCase{
		Repository: repo,
	}
}

func (u *GetCosmosUseCase) Execute(color string) ([]dto.GetCosmosResponse, error) {
	cosmos, err := u.Repository.GetCosmos(color)
	if err != nil {
		return nil, err
	}

	var result []dto.GetCosmosResponse
	for _, c := range cosmos {
		result = append(result, dto.GetCosmosResponse{
			ID:       c.ID,
			Name:     c.Name,
			Rank:     c.Rank,
			Color:    string(c.Color),
			ImageURL: c.ImageURL,
		})
	}

	return result, nil
}
