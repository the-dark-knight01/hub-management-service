package entity

type Team struct {
	ID    uint   `gorm:"primaryKey" json:"id,omitempty"`
	Name  string `gorm:"size:255;not null" json:"name" binding:"required,min=3,max=255"`
	HubID uint   `gorm:"not null" json:"hub_id" binding:"required"`
	Hub   *Hub   `gorm:"foreignKey:HubID;constraint:OnDelete:CASCADE" json:"hub,omitempty"`
}
