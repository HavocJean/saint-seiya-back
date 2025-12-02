package repositories

import (
	"saint-seiya-back/internal/domain/knight"
	"saint-seiya-back/internal/infrastructure/database"
)

type knightRepository struct{}

func NewKnightRepository() knight.Repository {
	return &knightRepository{}
}

func (r *knightRepository) GetKnights(page, limit int, rank, name string) ([]knight.Knight, error) {
	var result []knight.Knight
	query := database.DB.Model(&knight.Knight{})
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

func (r *knightRepository) GetKnightById(id uint) (*knight.Knight, error) {
	var result knight.Knight
	err := database.DB.First(&result, id).Error

	return &result, err
}

func (r *knightRepository) Create(k *knight.Knight) (*knight.Knight, error) {
	err := database.DB.Create(k).Error

	return k, err
}
