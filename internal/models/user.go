package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name 	  string    `gorm:"size:100;not null" json:"name"`
	Nickname  string    `gorm:"size:100;not null"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}