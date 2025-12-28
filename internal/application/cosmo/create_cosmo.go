package cosmo

import (
	"errors"
	"saint-seiya-back/internal/application/cosmo/dto"
	"saint-seiya-back/internal/domain/cosmo"
)

type CreateCosmoUseCase struct {
	Repository cosmo.Repository
}

func NewCreateCosmoUseCase(repo cosmo.Repository) *CreateCosmoUseCase {
	return &CreateCosmoUseCase{
		Repository: repo,
	}
}

type CreateCosmoInput struct {
	Name              string
	Rank              string
	Color             string
	SetBonusValue     float64
	SetBonusName      string
	SetBonusIsPercent bool
	ImageURL          *string
	BaseAttributes    []dto.CreateCosmoAttributeRequest
}

func (cc *CreateCosmoUseCase) Execute(input CreateCosmoInput) (*dto.CreateCosmoResponse, error) {
	validColors := map[string]bool{
		"red":       true,
		"blue":      true,
		"yellow":    true,
		"legendary": true,
	}

	if !validColors[input.Color] {
		return nil, errors.New("invalid cosmo color")
	}

	if len(input.BaseAttributes) == 0 {
		return nil, errors.New("base attributes are required")
	}

	baseAttributes := make([]cosmo.CosmoAttributeDomain, len(input.BaseAttributes))
	for i, attr := range input.BaseAttributes {
		baseAttributes[i] = cosmo.CosmoAttributeDomain{
			Name:      attr.Name,
			Value1:    attr.Value1,
			Value10:   attr.Value10,
			IsPercent: attr.IsPercent,
		}
	}

	cosmoDomain := &cosmo.CosmoDomain{
		Name:              input.Name,
		Rank:              input.Rank,
		Color:             cosmo.CosmoColor(input.Color),
		SetBonusValue:     input.SetBonusValue,
		SetBonusName:      input.SetBonusName,
		SetBonusIsPercent: input.SetBonusIsPercent,
		ImageURL:          input.ImageURL,
		BaseAttributes:    baseAttributes,
	}

	createdCosmo, err := cc.Repository.Create(cosmoDomain)
	if err != nil {
		return nil, errors.New("failed to create cosmo")
	}

	attributeResponses := make([]dto.CosmoAttributeResponse, len(createdCosmo.BaseAttributes))
	for i, attr := range createdCosmo.BaseAttributes {
		attributeResponses[i] = dto.CosmoAttributeResponse{
			ID:        attr.ID,
			CosmoID:   attr.CosmoID,
			Name:      attr.Name,
			Value1:    attr.Value1,
			Value10:   attr.Value10,
			IsPercent: attr.IsPercent,
		}
	}

	return &dto.CreateCosmoResponse{
		ID:                createdCosmo.ID,
		Name:              createdCosmo.Name,
		Rank:              createdCosmo.Rank,
		Color:             string(createdCosmo.Color),
		SetBonusValue:     createdCosmo.SetBonusValue,
		SetBonusName:      createdCosmo.SetBonusName,
		SetBonusIsPercent: createdCosmo.SetBonusIsPercent,
		BaseAttributes:    attributeResponses,
		ImageURL:          createdCosmo.ImageURL,
	}, nil
}
