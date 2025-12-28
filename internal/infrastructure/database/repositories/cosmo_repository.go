package repositories

import (
	"saint-seiya-back/internal/domain/cosmo"
	"saint-seiya-back/internal/infrastructure/database/entities"

	"gorm.io/gorm"
)

type cosmoRepository struct {
	db *gorm.DB
}

func NewCosmoRepository(db *gorm.DB) cosmo.Repository {
	return &cosmoRepository{db}
}

func (r *cosmoRepository) GetCosmos(color string) ([]cosmo.CosmoDomain, error) {
	var cosmos []entities.CosmoEntity

	query := r.db.Model(&entities.CosmoEntity{})

	if color != "" {
		query = query.Where("color = ?", color)
	}

	if err := query.Find(&cosmos).Error; err != nil {
		return nil, err
	}

	result := make([]cosmo.CosmoDomain, len(cosmos))
	for i, c := range cosmos {
		result[i] = cosmo.CosmoDomain{
			ID:                c.ID,
			Name:              c.Name,
			Rank:              c.Rank,
			Color:             cosmo.CosmoColor(c.Color),
			SetBonusValue:     c.SetBonusValue,
			SetBonusName:      c.SetBonusName,
			SetBonusIsPercent: c.SetBonusIsPercent,
			ImageURL:          c.ImageURL,
		}
	}

	return result, nil
}

func (r *cosmoRepository) GetCosmoByID(id uint) (*cosmo.CosmoDomain, error) {
	var cosmoEntity entities.CosmoEntity

	if err := r.db.Preload("BaseAttributes").First(&cosmoEntity, id).Error; err != nil {
		return nil, err
	}

	baseAttributes := make([]cosmo.CosmoAttributeDomain, len(cosmoEntity.BaseAttributes))
	for i, attr := range cosmoEntity.BaseAttributes {
		baseAttributes[i] = cosmo.CosmoAttributeDomain{
			ID:        attr.ID,
			CosmoID:   attr.CosmoID,
			Name:      attr.Name,
			Value1:    attr.Value1,
			Value10:   attr.Value10,
			IsPercent: attr.IsPercent,
		}
	}

	return &cosmo.CosmoDomain{
		ID:                cosmoEntity.ID,
		Name:              cosmoEntity.Name,
		Rank:              cosmoEntity.Rank,
		Color:             cosmo.CosmoColor(cosmoEntity.Color),
		SetBonusValue:     cosmoEntity.SetBonusValue,
		SetBonusName:      cosmoEntity.SetBonusName,
		SetBonusIsPercent: cosmoEntity.SetBonusIsPercent,
		BaseAttributes:    baseAttributes,
		ImageURL:          cosmoEntity.ImageURL,
	}, nil
}

func (r *cosmoRepository) Create(cosmoDomain *cosmo.CosmoDomain) (*cosmo.CosmoDomain, error) {
	cosmoEntity := &entities.CosmoEntity{
		Name:              cosmoDomain.Name,
		Rank:              cosmoDomain.Rank,
		Color:             string(cosmoDomain.Color),
		SetBonusValue:     cosmoDomain.SetBonusValue,
		SetBonusName:      cosmoDomain.SetBonusName,
		SetBonusIsPercent: cosmoDomain.SetBonusIsPercent,
		ImageURL:          cosmoDomain.ImageURL,
	}

	if err := r.db.Create(cosmoEntity).Error; err != nil {
		return nil, err
	}

	attributeEntities := make([]entities.CosmoAttributeEntity, len(cosmoDomain.BaseAttributes))
	for i, attr := range cosmoDomain.BaseAttributes {
		attributeEntities[i] = entities.CosmoAttributeEntity{
			CosmoID:   cosmoEntity.ID,
			Name:      attr.Name,
			Value1:    attr.Value1,
			Value10:   attr.Value10,
			IsPercent: attr.IsPercent,
		}
	}

	if err := r.db.Create(&attributeEntities).Error; err != nil {
		return nil, err
	}

	baseAttributes := make([]cosmo.CosmoAttributeDomain, len(attributeEntities))
	for i, attrEntity := range attributeEntities {
		baseAttributes[i] = cosmo.CosmoAttributeDomain{
			ID:        attrEntity.ID,
			CosmoID:   attrEntity.CosmoID,
			Name:      attrEntity.Name,
			Value1:    attrEntity.Value1,
			Value10:   attrEntity.Value10,
			IsPercent: attrEntity.IsPercent,
		}
	}

	return &cosmo.CosmoDomain{
		ID:                cosmoEntity.ID,
		Name:              cosmoEntity.Name,
		Rank:              cosmoEntity.Rank,
		Color:             cosmo.CosmoColor(cosmoEntity.Color),
		SetBonusValue:     cosmoEntity.SetBonusValue,
		SetBonusName:      cosmoEntity.SetBonusName,
		SetBonusIsPercent: cosmoEntity.SetBonusIsPercent,
		BaseAttributes:    baseAttributes,
		ImageURL:          cosmoEntity.ImageURL,
	}, nil
}
