package entities

import (
	"time"

	"gorm.io/gorm"
)

type CosmoEntity struct {
	ID                uint                   `gorm:"primaryKey" json:"id"`
	Name              string                 `gorm:"size:100" json:"name"`
	Rank              string                 `gorm:"size:10" json:"rank"`
	Color             string                 `gorm:"type:varchar(10)" json:"color"`
	SetBonusValue     float64                `gorm:"not null" json:"set_bonus"`
	SetBonusName      string                 `gorm:"size:50" json:"set_bonus_name"`
	SetBonusIsPercent bool                   `gorm:"not null" json:"set_bonus_is_percent"`
	BaseAttributes    []CosmoAttributeEntity `gorm:"foreignKey:CosmoID" json:"base_attributes"`
	ImageURL          *string                `gorm:"size:255" json:"image_url"`
	CreatedAt         time.Time              `json:"created_at"`
	UpdatedAt         time.Time              `json:"updated_at"`
	DeletedAt         gorm.DeletedAt         `gorm:"index" json:"-"`
}

func (CosmoEntity) TableName() string {
	return "cosmos"
}

type CosmoAttributeEntity struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CosmoID   uint           `gorm:"not null" json:"cosmo_id"`
	Name      string         `gorm:"size:100;not null" json:"name"`
	Value1    float64        `gorm:"not null" json:"value1"`
	Value10   float64        `gorm:"not null" json:"value_10"`
	IsPercent bool           `gorm:"not null" json:"is_percent"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (CosmoAttributeEntity) TableName() string {
	return "cosmo_attributes"
}
