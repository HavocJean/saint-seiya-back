package auth

import (
	"saint-seiya-back/internal/application/auth/dto"
	"saint-seiya-back/internal/domain/user"
)

type GetUserByIdUseCase struct {
	Repository user.Repository
}

func NewUserByIdUseCase(repo user.Repository) *GetUserByIdUseCase {
	return &GetUserByIdUseCase{Repository: repo}
}

func (u *GetUserByIdUseCase) Execute(id uint) (*dto.UserProfileResponse, error) {
	userDomain, err := u.Repository.GetUserById(id)

	if err != nil {
		return nil, err
	}

	return &dto.UserProfileResponse{
		Name:     userDomain.Name,
		Nickname: userDomain.Nickname,
		Email:    userDomain.Email,
	}, nil
}
