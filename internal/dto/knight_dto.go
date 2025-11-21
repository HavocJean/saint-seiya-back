package dto

import "saint-seiya-awakening/internal/models"

type CreateKnightRequest struct {
	Name         string  `json:"name" binding:"required"`
	Rank         string  `json:"rank" binding:"required"`
	Pv           int     `json:"pv" binding:"required"`
	AtkC         int     `json:"atk_c" binding:"required"`
	DefC         int     `json:"def_c" binding:"required"`
	DefF         int     `json:"def_f" binding:"required"`
	AtqF         int     `json:"atq_f" binding:"required"`
	Speed        int     `json:"speed" binding:"required"`
	StatusHit    float64 `json:"status_hit" binding:"required"`
	CritLevelF   float64 `json:"crit_level_f" binding:"required"`
	StatusResist float64 `json:"status_resist" binding:"required"`
	CritDamageC  float64 `json:"crit_damage_c" binding:"required"`
	CritEffectF  float64 `json:"crit_effect_f" binding:"required"`
	CritResistF  float64 `json:"crit_resist_f" binding:"required"`
	ImageURL     string  `json:"image_url"`
}

type CreateKnightSkillRequest struct {
	Name        string                    `json:"name" binding:"required"`
	Type        string                    `json:"type" binding:"required"`
	ImageURL    string                    `json:"image_url"`
	Description string                    `json:"description" binding:"required"`
	Levels      []models.KnightSkillLevel `json:"levels" binding:"required"`
}

type CreateKnightSkillLevel struct {
	Level       int    `json:"level" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type KnightsResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

// type SkillResponse struct {

// }
