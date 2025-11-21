package services

import (
	"fmt"
	"saint-seiya-awakening/internal/database"
	"saint-seiya-awakening/internal/dto"
	"saint-seiya-awakening/internal/models"
)

type KnightService struct{}

func NewKnightService() *KnightService {
	return &KnightService{}
}

func (s *KnightService) CreateKnight(req *dto.CreateKnightRequest) (*models.Knight, error) {

	if req.Speed <= 0 {
		return nil, fmt.Errorf("speed must be greater than zero")
	}

	knight := &models.Knight{
		Name:         req.Name,
		Rank:         req.Rank,
		Pv:           req.Pv,
		AtkC:         req.AtkC,
		DefC:         req.DefC,
		DefF:         req.DefF,
		AtqF:         req.AtqF,
		Speed:        req.Speed,
		StatusHit:    req.StatusHit,
		CritLevelF:   req.CritLevelF,
		StatusResist: req.StatusResist,
		CritDamageC:  req.CritDamageC,
		CritEffectF:  req.CritEffectF,
		CritResistF:  req.CritResistF,
		ImageURL:     req.ImageURL,
	}

	if err := database.DB.Create(knight).Error; err != nil {
		return nil, fmt.Errorf("failed to create knight: %w", err)
	}

	return knight, nil
}

func (s *KnightService) GetAllKnights(page, limit int, rank string) ([]dto.KnightsResponse, error) {
	var knights []dto.KnightsResponse
	offset := (page - 1) * limit

	query := database.DB.Model(&models.Knight{})

	if rank != "" {
		query = query.Where("rank = ?", rank)
	}

	if err := query.
		Offset(offset).
		Limit(limit).
		Find(&knights).Error; err != nil {
		return nil, fmt.Errorf("error to find knights: %w", err)
	}

	return knights, nil
}

func (s *KnightService) GetKnightByID(id uint) (*models.Knight, error) {
	var knight models.Knight

	if err := database.DB.
		Preload("Skills").
		First(&knight, id).Error; err != nil {
		return nil, fmt.Errorf("knight not found: %w", err)
	}

	return &knight, nil
}

func (s *KnightService) CreateKnightSkill(knightID uint, req *dto.CreateKnightSkillRequest) (*models.KnightSkill, error) {
	_, err := s.GetKnightByID((knightID))
	if err != nil {
		return nil, fmt.Errorf("knight not exists: %w", err)
	}

	if len(req.Levels) == 0 {
		return nil, fmt.Errorf("skill require level")
	}

	skillLevels := make([]models.KnightSkillLevel, len(req.Levels))
	for i, level := range req.Levels {
		skillLevels[i] = models.KnightSkillLevel{
			Level:       level.Level,
			Description: level.Description,
		}
	}

	skill := &models.KnightSkill{
		KnightID:    knightID,
		Name:        req.Name,
		Type:        req.Type,
		ImageURL:    req.ImageURL,
		Description: req.Description,
		Levels:      skillLevels,
	}

	if err := database.DB.Create(skill).Error; err != nil {
		return nil, fmt.Errorf("error to create skill: %w", err)
	}

	return skill, nil
}
