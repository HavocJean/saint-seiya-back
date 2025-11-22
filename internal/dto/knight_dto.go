package dto

import "saint-seiya-awakening/internal/models"

type CreateKnightRequest struct {
	Name  string `json:"name" binding:"required"`
	Rank  string `json:"rank" binding:"required"`
	Pv    int    `json:"pv" binding:"required"`
	AtkC  *int   `json:"atk_c" binding:"required"`
	DefC  int    `json:"def_c" binding:"required"`
	DefF  int    `json:"def_f" binding:"required"`
	AtkF  *int   `json:"atk_f" binding:"required"`
	Speed int    `json:"speed" binding:"required"`

	StatusHit       *float64 `json:"status_hit" binding:"required"`
	StatusResist    *float64 `json:"status_resist" binding:"required"`
	CritDamageC     *float64 `json:"crit_damage_c" binding:"required"`
	ResistDamageC   *float64 `json:"resist_damage_c" binding:"required"`
	PerfurationDefC *float64 `json:"perfuration_def_c"`
	ReflectDamage   *float64 `json:"reflect_damage" binding:"required"`
	Heal            *float64 `json:"heal" binding:"required"`

	CritLevelF      *float64 `json:"crit_level_f" binding:"required"`
	CritEffectF     *float64 `json:"crit_effect_f" binding:"required"`
	ResistCritF     *float64 `json:"resist_crit_f" binding:"required"`
	ResistDamageF   *float64 `json:"resist_damage_f" binding:"required"`
	PerfurationDefF *float64 `json:"perfuration_def_f" binding:"required"`
	LifeTheft       *float64 `json:"life_theft" binding:"required"`
	CritBasicF      *float64 `json:"crit_basic_f" binding:"required"`

	ImageURL string `json:"image_url"`
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
	Rank     string `json:"rank"`
	ImageUrl string `json:"image_url"`
}

// type SkillResponse struct {

// }
