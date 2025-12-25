package knight

import (
	"saint-seiya-back/internal/application/knight/dto"
	"saint-seiya-back/internal/domain/knight"
)

type GetKnightByIdUseCase struct {
	Repository knight.Repository
}

func NewGetKnightByIdUseCase(repo knight.Repository) *GetKnightByIdUseCase {
	return &GetKnightByIdUseCase{Repository: repo}
}

func (u *GetKnightByIdUseCase) Execute(id uint) (*dto.GetKnightByIdResponse, error) {
	knightDomain, err := u.Repository.GetKnightById(id)

	if err != nil {
		return nil, err
	}

	return &dto.GetKnightByIdResponse{
		ID:              knightDomain.ID,
		Name:            knightDomain.Name,
		Rank:            knightDomain.Rank,
		Pv:              knightDomain.Pv,
		AtkC:            knightDomain.AtkC,
		DefC:            knightDomain.DefC,
		DefF:            knightDomain.DefF,
		AtkF:            knightDomain.AtkF,
		Speed:           knightDomain.Speed,
		StatusHit:       knightDomain.StatusHit,
		StatusResist:    knightDomain.StatusResist,
		CritDamageC:     knightDomain.CritDamageC,
		ResistDamageC:   knightDomain.ResistDamageC,
		PerfurationDefC: knightDomain.PerfurationDefC,
		ReflectDamage:   knightDomain.ReflectDamage,
		Heal:            knightDomain.Heal,
		CritLevelF:      knightDomain.CritLevelF,
		CritEffectF:     knightDomain.CritEffectF,
		CritResistF:     knightDomain.CritResistF,
		ResistDamageF:   knightDomain.ResistDamageF,
		PerfurationDefF: knightDomain.PerfurationDefF,
		LifeTheft:       knightDomain.LifeTheft,
		CritBasicF:      knightDomain.CritBasicF,
		ImageURL:        knightDomain.ImageURL,
	}, nil
}
