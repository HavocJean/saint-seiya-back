package repositories

import (
	"saint-seiya-back/internal/domain/knight"
	"saint-seiya-back/internal/infrastructure/database"
)

type knightRepository struct{}

func NewKnightRepository() knight.Repository {
	return &knightRepository{}
}

func (r *knightRepository) GetKnights(page, limit int, rank, name string) ([]knight.KnightDomain, error) {
	var result []knight.KnightDomain
	query := database.DB.Model(&knight.KnightDomain{})
	offset := (page - 1) * limit

	if rank != "" {
		query = query.Where("rank = ?", rank)
	}

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	err := query.Offset(offset).Limit(limit).Find(&result).Error
	return result, err
}

func (r *knightRepository) GetKnightById(id uint) (*knight.KnightDomain, error) {
	var result knight.KnightDomain
	err := database.DB.First(&result, id).Error

	return &result, err
}

func (r *knightRepository) Create(k *knight.KnightDomain) (*knight.KnightDomain, error) {
	err := database.DB.Create(k).Error

	return k, err
}
