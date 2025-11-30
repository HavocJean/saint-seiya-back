package auth

import (
	"context"
	"errors"
	"fmt"
	"saint-seiya-back/internal/domain/user"
)

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	Token string
}

type LoginUseCase struct {
	UserRepository user.Repository
	JWTService     JWTService
}

func NewLoginUseCase(repo user.Repository, jwt JWTService) *LoginUseCase {
	return &LoginUseCase{
		UserRepository: repo,
		JWTService:     jwt,
	}
}

func (u *LoginUseCase) Execute(c context.Context, req LoginInput) (*LoginOutput, error) {
	userDomain, err := u.UserRepository.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !userDomain.VerifyPassword(req.Password) {
		return nil, errors.New("invalid credentials")
	}

	token, err := u.JWTService.GenerateToken(userDomain.ID, userDomain.Email)

	if err != nil {
		return nil, fmt.Errorf("failed to create token: %w", err)
	}

	return &LoginOutput{Token: token}, nil
}
