package repository

import (
	"gorm.io/gorm"
	"hub_management_service/internal/entity"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindUserByTeamID(teamID uint) ([]entity.User, error)
	FindByID(id uint) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

// Removed FindAll method, as per the request

// FindUserByTeamID - Method to find users by TeamID
func (r *userRepository) FindUserByTeamID(teamID uint) ([]entity.User, error) {
	var users []entity.User
	err := r.db.Where("team_id = ?", teamID).Find(&users).Error
	return users, err
}

// FindByID - Method to find a user by their ID
func (r *userRepository) FindByID(id uint) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
