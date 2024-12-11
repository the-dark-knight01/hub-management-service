package entity

type Hub struct {
	ID       uint    `gorm:"primaryKey" json:"id,omitempty"`
	Name     string  `gorm:"size:255;not null" json:"name" binding:"required,min=3,max=255"`
	Location string  `gorm:"size:255;not null" json:"location" binding:"required,min=3,max=255"`
	Teams    *[]Team `gorm:"foreignKey:HubID" json:"teams,omitempty"`
}
