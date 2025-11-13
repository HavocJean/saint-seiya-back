package models

import (
	"time"
)

type CosmoColor string

const (
	CosmoRed    CosmoColor = "red"
	CosmoBlue   CosmoColor = "blue"
	CosmoYellow CosmoColor = "yellow"
)

type Cosmo struct {
	ID                uint                `gorm:"primaryKey" json:"id"`
	Name              string              `gorm:"size:100" json:"name"`
	Rank              string              `gorm:"size:10" json:"rank"`
	Color             CosmoColor          `gorm:"type:varchar(10)" json:"color"`
	SetBonusValue     float64             `gorm:"not null" json:"set_bonus"`
	SetBonusName      string              `gorm:"size:50" json:"set_bonus_name"`
	SetBonusIsPercent bool                `gorm:"not null" json:"set_bonus_is_percent"`
	BaseAttributes    []CosmoAttribute    `gorm:"foreignKey:CosmoID" json:"base_attributes"`
	SubAttributes     []CosmoSubAttribute `gorm:"foreignKey:CosmoID" json:"sub_attributes"`
	ImageURL          *string             `gorm:"size:255" json:"image_url"`
	CreatedAt         time.Time           `json:"created_at"`
	UpdatedAt         time.Time           `json:"updated_at"`
}

type CosmoAttribute struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	CosmoID   uint    `gorm:"not null" json:"cosmo_id"`
	Name      string  `gorm:"size:100;not null" json:"name"`
	Value1    float64 `gorm:"not null" json:"value1"`
	Value10   float64 `gorm:"not null" json:"value_10"`
	IsPercent bool    `gorm:"not null" json:"is_percent"`
}

type CosmoSubAttribute struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	CosmoID   uint    `gorm:"not null" json:"cosmo_id"`
	Name      string  `gorm:"size:100;not null" json:"name"`
	Value     float64 `gorm:"not null" json:"value"`
	IsPercent bool    `gorm:"not null" json:"is_percent"`
}
