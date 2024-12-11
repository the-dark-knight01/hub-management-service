package repository

import (
	"gorm.io/gorm"
	"hub_management_service/internal/entity"
)

type HubRepository interface {
	Create(hub *entity.Hub) error
	FindAll() ([]entity.Hub, error)
	FindByID(id uint) (*entity.Hub, error)
	SearchByName(name string) ([]entity.Hub, error)
}

type hubRepository struct {
	db *gorm.DB
}

func NewHubRepository(db *gorm.DB) HubRepository {
	return &hubRepository{db: db}
}

func (r *hubRepository) Create(hub *entity.Hub) error {
	return r.db.Create(hub).Error
}

func (r *hubRepository) FindAll() ([]entity.Hub, error) {
	var hubs []entity.Hub
	err := r.db.Find(&hubs).Error
	return hubs, err
}

func (r *hubRepository) FindByID(id uint) (*entity.Hub, error) {
	var hub entity.Hub
	// Use First to find a record by ID
	err := r.db.First(&hub, id).Error
	if err != nil {
		return nil, err
	}
	return &hub, nil
}

func (r *hubRepository) SearchByName(name string) ([]entity.Hub, error) {
	var hubs []entity.Hub
	// Preload related teams and search for hubs by name
	err := r.db.Preload("Teams").Where("name LIKE ?", "%"+name+"%").Find(&hubs).Error
	return hubs, err
}
