package models

import (
	"time"
)

type Team struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"size:100;not null" json:"name"`
	UserID    uint       `gorm:"not null" json:"user_id"`
	User      User       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"user,omitempty"`
	IsPublic  bool       `gorm:"not null" json:"is_public"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

type TeamKnight struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	TeamID    uint       `gorm:"not null" json:"team_id"`
	Team      Team       `gorm:"foreignKey:TeamID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"team,omitempty"`
	KnightID  uint       `gorm:"not null" json:"knight_id"`
	Knight    Knight     `gorm:"foreignKey:KnightID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"knight,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

type TeamVote struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	TeamID    uint       `gorm:"not null;uniqueIndex:uniq_team_user" json:"team_id"`
	Team      Team       `gorm:"foreignKey:TeamID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"team,omitempty"`
	UserID    uint       `gorm:"not null;uniqueIndex:uniq_team_user" json:"user_id"`
	User      User       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	Vote      int16      `gorm:"not null" json:"vote"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

type TeamFavorite struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	TeamID    uint       `gorm:"not null;uniqueIndex:uniq_team_user" json:"team_id"`
	Team      Team       `gorm:"foreignKey:TeamID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"team,omitempty"`
	UserID    uint       `gorm:"not null;uniqueIndex:uniq_team_user" json:"user_id"`
	User      User       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
