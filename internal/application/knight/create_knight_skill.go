package knight

import (
	"errors"
	"saint-seiya-back/internal/application/knight/dto"
	"saint-seiya-back/internal/domain/knight"
)

type CreateKnightSkillUseCase struct {
	knightRepository knight.Repository
	skillRepository  knight.Repository
}

func NewCreateKnightSkillUseCase(repo knight.Repository) *CreateKnightSkillUseCase {
	return &CreateKnightSkillUseCase{
		knightRepository: repo,
		skillRepository:  repo,
	}
}

type CreateKnightSkillInput struct {
	KnightID    uint
	Name        string
	Type        string
	ImageURL    *string
	Description string
	Levels      []dto.CreateKnightSkillLevelRequest
}

func (ck *CreateKnightSkillUseCase) Execute(input CreateKnightSkillInput) (*dto.CreateKnightSkillResponse, error) {
	_, err := ck.knightRepository.GetKnightById(input.KnightID)
	if err != nil {
		return nil, errors.New("knight not exists")
	}

	if len(input.Levels) == 0 {
		return nil, errors.New("skill require level")
	}

	levels := make([]knight.KnightSkillLevelDomain, len(input.Levels))
	for i, level := range input.Levels {
		levels[i] = knight.KnightSkillLevelDomain{
			Level:       level.Level,
			Description: level.Description,
		}
	}

	skillDomain := &knight.KnightSkillDomain{
		KnightID:    input.KnightID,
		Name:        input.Name,
		Type:        input.Type,
		ImageURL:    input.ImageURL,
		Description: input.Description,
		Levels:      levels,
	}

	createdSkill, err := ck.skillRepository.CreateSkill(skillDomain)
	if err != nil {
		return nil, errors.New("failed to create skill")
	}

	levelResponse := make([]dto.CreateKnightSkillLevelResponse, len(createdSkill.Levels))
	for i, level := range createdSkill.Levels {
		levelResponse[i] = dto.CreateKnightSkillLevelResponse{
			ID:          level.ID,
			SkillID:     level.SkillID,
			Level:       level.Level,
			Description: level.Description,
		}
	}

	return &dto.CreateKnightSkillResponse{
		ID:          createdSkill.ID,
		KnightID:    createdSkill.KnightID,
		Name:        createdSkill.Name,
		Type:        createdSkill.Type,
		ImageURL:    createdSkill.ImageURL,
		Description: createdSkill.Description,
		Levels:      levelResponse,
	}, nil
}
