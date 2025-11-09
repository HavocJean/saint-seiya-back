package models

import (
	"time"
)

type Cosmo struct {
	ID             uint                `gorm:"primaryKey" json:"id"`
	Name           string              `gorm:"size:100" json:"name"`
	Rank           string              `gorm:"size:10" json:"rank"`
	SetBonus       string              `gorm:"size:50" json:"set_bonus"`
	BaseAttributes []CosmoAttribute    `gorm:"foreignKey:CosmoID" json:"base_attributes"`
	SubAttributes  []CosmoSubAttribute `gorm:"foreignKey:CosmoID" json:"sub_attributes"`
	ImageURL       *string             `gorm:"size:255" json:"image_url"`
	CreatedAt      time.Time           `json:"created_at"`
	UpdatedAt      time.Time           `json:"updated_at"`
}

type CosmoAttribute struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	CosmoID     uint    `gorm:"not null" json:"cosmo_id"`
	Name        string  `gorm:"size:100;not null" json:"name"`
	NameValue1  string  `gorm:"size:50;not null" json:"name_value1"`
	Value1      float64 `gorm:"not null" json:"value1"`
	NameValue10 string  `gorm:"size:50" json:"name_value_10"`
	Value10     float64 `json:"value_10"`
}

type CosmoSubAttribute struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	CosmoID   uint    `gorm:"not null" json:"cosmo_id"`
	Name      string  `gorm:"size:100;not null" json:"name"`
	Value     float64 `gorm:"not null" json:"value"`
	IsPercent bool    `gorm:"not null" json:"is_percent"`
}
