package entities

import (
	"time"

	"gorm.io/gorm"
)

type KnightEntity struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Name            string         `gorm:"size:100;not null" json:"name"`
	Rank            string         `gorm:"size:10;not null" json:"rank"`
	Pv              int            `gorm:"not null" json:"pv"`
	AtkC            *int           `json:"atk_c"`
	DefC            int            `gorm:"not null" json:"def_c"`
	DefF            int            `gorm:"not null" json:"def_f"`
	AtkF            *int           `json:"atk_f"`
	Speed           int            `gorm:"not null" json:"speed"`
	StatusHit       *float64       `json:"status_hit"`
	CritLevelF      *float64       `json:"crit_level_f"`
	StatusResist    *float64       `json:"status_resist"`
	CritDamageC     *float64       `json:"crit_damage_c"`
	ResistDamageC   *float64       `json:"resist_damage_c"`
	PerfurationDefC *float64       `json:"perfuration_def_c"`
	ReflectDamage   *float64       `json:"reflect_damage"`
	Heal            *float64       `json:"heal"`
	CritEffectF     *float64       `json:"crit_effect_f"`
	CritResistF     *float64       `json:"crit_resist_f"`
	ResistDamageF   *float64       `json:"resist_damage_f"`
	PerfurationDefF *float64       `json:"perfuration_def_f"`
	LifeTheft       *float64       `json:"life_theft"`
	CritBasicF      *float64       `json:"crit_basic_f"`
	ImageURL        *string        `gorm:"size:255" json:"image_url"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

func (KnightEntity) TableName() string {
	return "knights"
}
