package entity

type User struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"size:255;not null" json:"name" binding:"required"`
	TeamID uint   `gorm:"not null" json:"team_id" binding:"required"`
	Email  string `gorm:"not null" json:"email" binding:"required"`
	Team   *Team  `gorm:"foreignKey:TeamID;constraint:OnDelete:CASCADE" json:"team,omitempty"`
}
