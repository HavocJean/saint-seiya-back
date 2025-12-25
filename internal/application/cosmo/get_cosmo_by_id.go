package cosmo

import (
	"saint-seiya-back/internal/application/cosmo/dto"
	"saint-seiya-back/internal/domain/cosmo"
)

type GetCosmoByIdUseCase struct {
	Repository cosmo.Repository
}

func NewGetCosmoByIdUseCase(repo cosmo.Repository) *GetCosmoByIdUseCase {
	return &GetCosmoByIdUseCase{
		Repository: repo,
	}
}

func (u *GetCosmoByIdUseCase) Execute(id uint) (*dto.GetCosmoByIdResponse, error) {
	cosmo, err := u.Repository.GetCosmoByID(id)
	if err != nil {
		return nil, err
	}

	var baseAttributes []dto.CosmoAttributeResponse
	for _, attr := range cosmo.BaseAttributes {
		baseAttributes = append(baseAttributes, dto.CosmoAttributeResponse{
			ID:        attr.ID,
			CosmoID:   attr.CosmoID,
			Name:      attr.Name,
			Value1:    attr.Value1,
			Value10:   attr.Value10,
			IsPercent: attr.IsPercent,
		})
	}

	return &dto.GetCosmoByIdResponse{
		ID:                cosmo.ID,
		Name:              cosmo.Name,
		Rank:              cosmo.Rank,
		Color:             string(cosmo.Color),
		SetBonusValue:     cosmo.SetBonusValue,
		SetBonusName:      cosmo.SetBonusName,
		SetBonusIsPercent: cosmo.SetBonusIsPercent,
		BaseAttributes:    baseAttributes,
		ImageURL:          cosmo.ImageURL,
	}, nil
}
