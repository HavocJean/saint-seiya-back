package auth

import (
	"context"
	"errors"
	"saint-seiya-back/internal/domain/user"
)

type RegisterInput struct {
	Name     string
	Nickname string
	Email    string
	Password string
}

type RegisterOutput struct {
	Token string
}

type RegisterUseCase struct {
	UserRepository user.Repository
	JWTService     JWTService
}

func NewRegisterUseCase(repo user.Repository, jwt JWTService) *RegisterUseCase {
	return &RegisterUseCase{
		UserRepository: repo,
		JWTService:     jwt,
	}
}

func (u *RegisterUseCase) Execute(c context.Context, req RegisterInput) (*RegisterOutput, error) {
	_, err := u.UserRepository.FindByEmail(req.Email)
	if err == nil {
		return nil, errors.New("email already exists")
	}

	newUser := user.NewUser(
		req.Name,
		req.Nickname,
		req.Email,
		req.Password,
	)

	if err := u.UserRepository.Create(newUser); err != nil {
		return nil, err
	}

	token, err := u.JWTService.GenerateToken(newUser.ID, newUser.Email)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &RegisterOutput{Token: token}, nil
}
