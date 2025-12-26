package repositories

import (
	"saint-seiya-back/internal/domain/knight"
	"saint-seiya-back/internal/infrastructure/database/entities"

	"gorm.io/gorm"
)

type knightRepository struct {
	db *gorm.DB
}

func NewKnightRepository(db *gorm.DB) knight.Repository {
	return &knightRepository{db}
}

func (r *knightRepository) GetKnights(page, limit int, rank, name string) ([]knight.KnightDomain, error) {
	var knightsEntities []entities.KnightEntity
	offset := (page - 1) * limit

	query := r.db.Model(&entities.KnightEntity{})

	if rank != "" {
		query = query.Where("rank = ?", rank)
	}

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	if err := query.Offset(offset).Limit(limit).Find(&knightsEntities).Error; err != nil {
		return nil, err
	}

	result := make([]knight.KnightDomain, len(knightsEntities))
	for i, k := range knightsEntities {
		result[i] = knight.KnightDomain{
			ID:              k.ID,
			Name:            k.Name,
			Rank:            k.Rank,
			Pv:              k.Pv,
			AtkC:            k.AtkC,
			DefC:            k.DefC,
			DefF:            k.DefF,
			AtkF:            k.AtkF,
			Speed:           k.Speed,
			StatusHit:       k.StatusHit,
			CritLevelF:      k.CritLevelF,
			StatusResist:    k.StatusResist,
			CritDamageC:     k.CritDamageC,
			ResistDamageC:   k.ResistDamageC,
			PerfurationDefC: k.PerfurationDefC,
			ReflectDamage:   k.ReflectDamage,
			Heal:            k.Heal,
			CritEffectF:     k.CritEffectF,
			CritResistF:     k.CritResistF,
			ResistDamageF:   k.ResistDamageF,
			PerfurationDefF: k.PerfurationDefF,
			LifeTheft:       k.LifeTheft,
			CritBasicF:      k.CritBasicF,
			ImageURL:        k.ImageURL,
		}
	}

	return result, nil
}

func (r *knightRepository) GetKnightById(id uint) (*knight.KnightDomain, error) {
	var knightEntity entities.KnightEntity

	if err := r.db.First(&knightEntity, id).Error; err != nil {
		return nil, err
	}

	return &knight.KnightDomain{
		ID:              knightEntity.ID,
		Name:            knightEntity.Name,
		Rank:            knightEntity.Rank,
		Pv:              knightEntity.Pv,
		AtkC:            knightEntity.AtkC,
		DefC:            knightEntity.DefC,
		DefF:            knightEntity.DefF,
		AtkF:            knightEntity.AtkF,
		Speed:           knightEntity.Speed,
		StatusHit:       knightEntity.StatusHit,
		CritLevelF:      knightEntity.CritLevelF,
		StatusResist:    knightEntity.StatusResist,
		CritDamageC:     knightEntity.CritDamageC,
		ResistDamageC:   knightEntity.ResistDamageC,
		PerfurationDefC: knightEntity.PerfurationDefC,
		ReflectDamage:   knightEntity.ReflectDamage,
		Heal:            knightEntity.Heal,
		CritEffectF:     knightEntity.CritEffectF,
		CritResistF:     knightEntity.CritResistF,
		ResistDamageF:   knightEntity.ResistDamageF,
		PerfurationDefF: knightEntity.PerfurationDefF,
		LifeTheft:       knightEntity.LifeTheft,
		CritBasicF:      knightEntity.CritBasicF,
		ImageURL:        knightEntity.ImageURL,
	}, nil
}

func (r *knightRepository) Create(k *knight.KnightDomain) (*knight.KnightDomain, error) {
	knightEntity := &entities.KnightEntity{
		Name:            k.Name,
		Rank:            k.Rank,
		Pv:              k.Pv,
		AtkC:            k.AtkC,
		DefC:            k.DefC,
		DefF:            k.DefF,
		AtkF:            k.AtkF,
		Speed:           k.Speed,
		StatusHit:       k.StatusHit,
		CritLevelF:      k.CritLevelF,
		StatusResist:    k.StatusResist,
		CritDamageC:     k.CritDamageC,
		ResistDamageC:   k.ResistDamageC,
		PerfurationDefC: k.PerfurationDefC,
		ReflectDamage:   k.ReflectDamage,
		Heal:            k.Heal,
		CritEffectF:     k.CritEffectF,
		CritResistF:     k.CritResistF,
		ResistDamageF:   k.ResistDamageF,
		PerfurationDefF: k.PerfurationDefF,
		LifeTheft:       k.LifeTheft,
		CritBasicF:      k.CritBasicF,
		ImageURL:        k.ImageURL,
	}

	if err := r.db.Create(knightEntity).Error; err != nil {
		return nil, err
	}

	return &knight.KnightDomain{
		ID:              knightEntity.ID,
		Name:            knightEntity.Name,
		Rank:            knightEntity.Rank,
		Pv:              knightEntity.Pv,
		AtkC:            knightEntity.AtkC,
		DefC:            knightEntity.DefC,
		DefF:            knightEntity.DefF,
		AtkF:            knightEntity.AtkF,
		Speed:           knightEntity.Speed,
		StatusHit:       knightEntity.StatusHit,
		CritLevelF:      knightEntity.CritLevelF,
		StatusResist:    knightEntity.StatusResist,
		CritDamageC:     knightEntity.CritDamageC,
		ResistDamageC:   knightEntity.ResistDamageC,
		PerfurationDefC: knightEntity.PerfurationDefC,
		ReflectDamage:   knightEntity.ReflectDamage,
		Heal:            knightEntity.Heal,
		CritEffectF:     knightEntity.CritEffectF,
		CritResistF:     knightEntity.CritResistF,
		ResistDamageF:   knightEntity.ResistDamageF,
		PerfurationDefF: knightEntity.PerfurationDefF,
		LifeTheft:       knightEntity.LifeTheft,
		CritBasicF:      knightEntity.CritBasicF,
		ImageURL:        knightEntity.ImageURL,
	}, nil
}
