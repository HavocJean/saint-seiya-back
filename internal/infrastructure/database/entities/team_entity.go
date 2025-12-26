package entities

import (
	"time"

	"gorm.io/gorm"
)

type TeamEntity struct {
	ID          uint               `gorm:"primaryKey" json:"id"`
	Name        string             `gorm:"size:100;not null" json:"name"`
	UserID      uint               `gorm:"not null" json:"user_id"`
	IsPublic    bool               `gorm:"not null" json:"is_public"`
	TeamKnights []TeamKnightEntity `gorm:"foreignKey:TeamID" json:"team_knights,omitempty"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	DeletedAt   gorm.DeletedAt     `gorm:"index" json:"-"`
}

func (TeamEntity) TableName() string {
	return "teams"
}

type TeamKnightEntity struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	TeamID    uint           `gorm:"not null" json:"team_id"`
	KnightID  uint           `gorm:"not null" json:"knight_id"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (TeamKnightEntity) TableName() string {
	return "team_knights"
}
