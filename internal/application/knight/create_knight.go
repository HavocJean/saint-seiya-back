package knight

import (
	"errors"
	"saint-seiya-back/internal/domain/knight"
)

type CreateKnightUseCase struct {
	repo knight.Repository
}

func NewCreateKnightUseCase(r knight.Repository) *CreateKnightUseCase {
	return &CreateKnightUseCase{repo: r}
}

func (uc *CreateKnightUseCase) Execute(k *knight.KnightDomain) (*knight.KnightDomain, error) {

	if k.Name == "" {
		return nil, errors.New("name is required")
	}

	if k.Rank == "" {
		return nil, errors.New("rank is required")
	}

	if k.Speed <= 0 {
		return nil, errors.New("speed must be greater than zero")
	}

	return uc.repo.Create(k)
}
