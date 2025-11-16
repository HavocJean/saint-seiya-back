package models

import (
	"time"
)

type Knight struct {
	ID           uint          `gorm:"primaryKey" json:"id"`
	Name         string        `gorm:"size:100;not null" json:"name"`
	Rank         string        `gorm:"size:10;not null" json:"rank"`
	Pv           int           `gorm:"not null" json:"pv"`
	AtkC         int           `gorm:"not null" json:"atk_c"`
	DefC         int           `gorm:"not null" json:"def_c"`
	DefF         int           `gorm:"not null" json:"def_f"`
	AtqF         int           `gorm:"not null" json:"atq_f"`
	Speed        int           `gorm:"not null" json:"speed"`
	StatusHit    float64       `gorm:"not null" json:"status_hit"`
	CritLevelF   float64       `gorm:"not null" json:"crit_level_f"`
	StatusResist float64       `gorm:"not null" json:"status_resist"`
	CritDamageC  float64       `gorm:"not null" json:"crit_damage_c"`
	CritEffectF  float64       `gorm:"not null" json:"crit_effect_f"`
	CritResistF  float64       `gorm:"not null" json:"crit_resist_f"`
	ImageURL     string        `gorm:"size:255" json:"image_url"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	Skills       []KnightSkill `gorm:"foreignKey:KnightID" json:"skills"`
}

type KnightSkill struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	KnightID    uint   `gorm:"not null"`
	Name        string `gorm:"size:150;not null"`
	Type        string `gorm:"size:50"`
	ImageURL    string `gorm:"size:255" json:"image_url"`
	Description string
	Levels      []KnightSkillLevel `gorm:"foreignKey:SkillID" json:"levels"`
}

type KnightSkillLevel struct {
	ID          uint `gorm:"primaryKey" json:"id"`
	SkillID     uint `gorm:"not null"`
	Level       int  `gorm:"not null"`
	Description string
}
