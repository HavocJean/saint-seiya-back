package entities

import (
	"saint-seiya-back/internal/domain/user"
	"time"

	"gorm.io/gorm"
)

type UserEntity struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:100;not null" json:"name"`
	Nickname  string         `gorm:"size:100;not null" json:"nickname"`
	Email     string         `gorm:"unique;not null" json:"email"`
	Password  string         `gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (u *UserEntity) BeforeCreate(tx *gorm.DB) error {
	if u.Password != "" {
		hashedPassword, err := user.HashedPassword(u.Password)

		if err != nil {
			return err
		}
		u.Password = hashedPassword
	}
	return nil
}
