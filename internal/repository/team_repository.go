package repository

import (
	"gorm.io/gorm"
	"hub_management_service/internal/entity"
)

type TeamRepository interface {
	Create(team *entity.Team) error
	FindAll() ([]entity.Team, error)
	FindByHubID(hubID uint) ([]entity.Team, error)
	FindByID(id uint) (*entity.Team, error)
}

type teamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) TeamRepository {
	return &teamRepository{db: db}
}

func (r *teamRepository) Create(team *entity.Team) error {
	return r.db.Create(team).Error
}

func (r *teamRepository) FindAll() ([]entity.Team, error) {
	var teams []entity.Team
	err := r.db.Find(&teams).Error
	return teams, err
}

func (r *teamRepository) FindByHubID(hubID uint) ([]entity.Team, error) {
	var teams []entity.Team
	err := r.db.Where("hub_id = ?", hubID).Find(&teams).Error
	return teams, err
}

func (r *teamRepository) FindByID(id uint) (*entity.Team, error) {
	var team entity.Team
	err := r.db.First(&team, id).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}
