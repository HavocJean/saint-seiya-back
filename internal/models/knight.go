package models

import (
	"time"
)

type Knight struct {
	ID				uint	  `gorm:"primaryKey" json:"id"`
	Name			string    `gorm:"size:100;not null" json:"name"`
	Rank			string    `gorm:"size:50;not null" json:"rank"`
	Pv				int       `gorm:"not null" json:"pv"`
	AtkC			int       `gorm:"not null" json:"atk_c"`
	DefC			int       `gorm:"not null" json:"def_c"`
	DefF			int       `gorm:"not null" json:"def_f"`
	AtqF			int       `gorm:"not null" json:"atq_f"`
	Speed			int       `gorm:"not null" json:"speed"`
	StatusHit		float64   `gorm:"not null" json:"status_hit"`
	CritLevelF		float64   `gorm:"not null" json:"crit_level_f"`
	StatusResist	float64   `gorm:"not null" json:"status_resist"`
	CritDamageC		float64   `gorm:"not null" json:"crit_damage_c"`
	CritEffectF		float64   `gorm:"not null" json:"crit_effect_f"`
	CritResistF     float64   `gorm:"not null" json:"crit_resist_f"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}