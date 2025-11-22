package models

import (
	"time"
)

type Knight struct {
	ID              uint          `gorm:"primaryKey" json:"id"`
	Name            string        `gorm:"size:100;not null" json:"name"`
	Rank            string        `gorm:"size:10;not null" json:"rank"`
	Pv              int           `gorm:"not null" json:"pv"`
	AtkC            int           `gorm:"not null" json:"atk_c"`
	DefC            int           `gorm:"not null" json:"def_c"`
	DefF            int           `gorm:"not null" json:"def_f"`
	AtkF            int           `gorm:"not null" json:"atk_f"`
	Speed           int           `gorm:"not null" json:"speed"`
	StatusHit       float64       `gorm:"not null" json:"status_hit"`
	CritLevelF      float64       `gorm:"not null" json:"crit_level_f"`
	StatusResist    float64       `gorm:"not null" json:"status_resist"`
	CritDamageC     float64       `gorm:"not null" json:"crit_damage_c"`
	ResistDamageC   float64       `gorm:"not null" json:"resist_damage_c"`
	PerfurationDefC float64       `gorm:"not null" json:"perfuration_def_c"`
	ReflectDamage   float64       `gorm:"not null" json:"reflect_damage"`
	Heal            float64       `gorm:"not null" json:"heal"`
	CritEffectF     float64       `gorm:"not null" json:"crit_effect_f"`
	CritResistF     float64       `gorm:"not null" json:"crit_resist_f"`
	ResistDamageF   float64       `gorm:"not null" json:"resist_damage_f"`
	PerfurationDefF float64       `gorm:"not null" json:"perfuration_def_f"`
	LifeTheft       float64       `gorm:"not null" json:"life_theft"`
	CritBasicF      float64       `gorm:"not null" json:"crit_basic_f"`
	ImageURL        string        `gorm:"size:255" json:"image_url"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	Skills          []KnightSkill `gorm:"foreignKey:KnightID" json:"skills"`
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
