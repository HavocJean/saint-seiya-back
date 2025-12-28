package repositories

import (
	"saint-seiya-back/internal/domain/user"
	"saint-seiya-back/internal/infrastructure/database/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindByEmail(email string) (*user.User, error) {
	var u entities.UserEntity
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}

	return &user.User{
		ID:       u.ID,
		Name:     u.Name,
		Nickname: u.Nickname,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}

func (r *UserRepository) Create(u *user.User) error {
	entity := &entities.UserEntity{
		Name:     u.Name,
		Nickname: u.Nickname,
		Email:    u.Email,
		Password: u.Password,
	}

	return r.db.Create(entity).Error
}

func (r *UserRepository) GetUserById(id uint) (*user.User, error) {
	var u entities.UserEntity
	if err := r.db.First(&u, id).Error; err != nil {
		return nil, err
	}

	return &user.User{
		ID:       u.ID,
		Name:     u.Name,
		Nickname: u.Nickname,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}
