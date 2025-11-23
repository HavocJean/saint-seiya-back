package models

import (
	"time"
)

type Team struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"size:100;not null" json:"name"`
	UserID    uint       `gorm:"not null" json:"user_id"`
	IsPublic  bool       `gorm:"not null" json:"is_public"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

type TeamKnight struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	TeamID    uint       `gorm:"not null" json:"team_id"`
	KnightID  uint       `gorm:"not null" json:"knight_id"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

type TeamVote struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	TeamID    uint      `gorm:"not null;uniqueIndex:uniq_team_user" json:"team_id"`
	UserID    uint      `gorm:"not null;uniqueIndex:uniq_team_user" json:"user_id"`
	Vote      int16     `gorm:"not null" json:"vote"`
	CreatedAt time.Time `json:"created_at"`
}

type TeamFavorite struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	TeamID    uint       `gorm:"not null;uniqueIndex:uniq_team_user" json:"team_id"`
	UserID    uint       `gorm:"not null;uniqueIndex:uniq_team_user" json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
