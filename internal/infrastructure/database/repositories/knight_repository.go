package repositories

import (
	"saint-seiya-back/internal/domain/knight"

	"gorm.io/gorm"
)

type knightRepository struct {
	db *gorm.DB
}

func NewKnightRepository(db *gorm.DB) knight.Repository {
	return &knightRepository{db}
}

func (r *knightRepository) GetKnights(page, limit int, rank, name string) ([]knight.KnightDomain, error) {
	var result []knight.KnightDomain
	query := r.db.Model(&knight.KnightDomain{})
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
	err := r.db.First(&result, id).Error

	return &result, err
}

func (r *knightRepository) Create(k *knight.KnightDomain) (*knight.KnightDomain, error) {
	err := r.db.Create(k).Error

	return k, err
}
