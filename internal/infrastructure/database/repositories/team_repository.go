package repositories

import (
	"saint-seiya-back/internal/domain/team"
	"saint-seiya-back/internal/infrastructure/database/entities"

	"gorm.io/gorm"
)

type teamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) team.Repository {
	return &teamRepository{db: db}
}

func (r *teamRepository) Create(teamDomain *team.TeamDomain) (*team.TeamDomain, error) {
	teamEntity := &entities.TeamEntity{
		Name:     teamDomain.Name,
		UserID:   teamDomain.UserID,
		IsPublic: teamDomain.IsPublic,
	}

	if err := r.db.Create(teamEntity).Error; err != nil {
		return nil, err
	}

	return &team.TeamDomain{
		ID:       teamEntity.ID,
		Name:     teamEntity.Name,
		UserID:   teamEntity.UserID,
		IsPublic: teamEntity.IsPublic,
	}, nil
}

func (r *teamRepository) GetByID(id uint) (*team.TeamDomain, error) {
	var teamEntity entities.TeamEntity

	if err := r.db.First(&teamEntity, id).Error; err != nil {
		return nil, err
	}

	return &team.TeamDomain{
		ID:       teamEntity.ID,
		Name:     teamEntity.Name,
		UserID:   teamEntity.UserID,
		IsPublic: teamEntity.IsPublic,
	}, nil
}

func (r *teamRepository) GetByUserID(userID uint) ([]team.TeamDomain, error) {
	var teamsEntities []entities.TeamEntity

	if err := r.db.Where("user_id = ?", userID).Find(&teamsEntities).Error; err != nil {
		return nil, err
	}

	result := make([]team.TeamDomain, len(teamsEntities))
	for i, t := range teamsEntities {
		result[i] = team.TeamDomain{
			ID:       t.ID,
			Name:     t.Name,
			UserID:   t.UserID,
			IsPublic: t.IsPublic,
		}
	}

	return result, nil
}

func (r *teamRepository) CountByUserID(userID uint) (int64, error) {
	var count int64
	err := r.db.Model(&entities.TeamEntity{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}

func (r *teamRepository) Delete(id uint, userID uint) error {
	if err := r.db.Where("team_id = ?", id).Delete(&entities.TeamKnightEntity{}).Error; err != nil {
		return err
	}

	if err := r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&entities.TeamEntity{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *teamRepository) AddKnightToTeam(teamKnightDomain *team.TeamKnightDomain) (*team.TeamKnightDomain, error) {
	teamKnightEntity := &entities.TeamKnightEntity{
		TeamID:   teamKnightDomain.TeamID,
		KnightID: teamKnightDomain.KnightID,
	}

	if err := r.db.Create(teamKnightEntity).Error; err != nil {
		return nil, err
	}

	return &team.TeamKnightDomain{
		ID:       teamKnightEntity.ID,
		TeamID:   teamKnightEntity.TeamID,
		KnightID: teamKnightEntity.KnightID,
	}, nil
}

func (r *teamRepository) CountKnightsByTeamID(teamID uint) (int64, error) {
	var count int64
	err := r.db.Model(&entities.TeamKnightEntity{}).Where("team_id = ?", teamID).Count(&count).Error
	return count, err
}

func (r *teamRepository) KnightExistsInTeam(teamID uint, knightID uint) (bool, error) {
	var count int64
	err := r.db.Model(&entities.TeamKnightEntity{}).
		Where("team_id = ? AND knight_id = ?", teamID, knightID).
		Count(&count).Error

	return count > 0, err
}

func (r *teamRepository) DeleteKnightToTeam(teamID uint, knightID uint) error {
	if err := r.db.Where("team_id = ? AND knight_id = ?", teamID, knightID).
		Delete(&entities.TeamKnightEntity{}).Error; err != nil {
		return err
	}

	return nil
}
