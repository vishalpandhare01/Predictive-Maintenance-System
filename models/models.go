package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Equipment struct for the equipment table
type Equipment struct {
	ID        string    `gorm:"type:char(36);primarykey"`
	Name      string    `gorm:"type:varchar(255)"`
	Type      string    `gorm:"type:varchar(50)"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}

func (U *Equipment) BeforeCreate(tx *gorm.DB) (err error) {
	U.ID = uuid.New().String()
	return
}

// Sensor struct for the sensors table
type Sensor struct {
	ID          string    `gorm:"type:char(36);primarykey"`
	EquipmentID string    `gorm:"type:varchar(36);not:null"`
	Type        string    `gorm:"type:varchar(50)"`
	Value       float64   `gorm:"type:float"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
	Equipment   Equipment `gorm:"foreignKey:EquipmentID"`
}

func (U *Sensor) BeforeCreate(tx *gorm.DB) (err error) {
	U.ID = uuid.New().String()
	return
}

// MaintenanceLog struct for the maintenance_logs table
type MaintenanceLog struct {
	ID          string    `gorm:"type:char(36);primarykey"`
	EquipmentID string    `gorm:"type:varchar(36);not:null"`
	Date        time.Time `gorm:"autoCreateTime"` // Automatically set the current timestamp
	Description string    `gorm:"type:text"`
	Equipment   Equipment `gorm:"foreignKey:EquipmentID"`
}

func (U *MaintenanceLog) BeforeCreate(tx *gorm.DB) (err error) {
	U.ID = uuid.New().String()
	return
}

// Prediction struct for the predictions table
type Prediction struct {
	ID                   string    `gorm:"type:char(36);primarykey"`
	EquipmentID          string    `gorm:"type:varchar(36);not:null"`
	PredictedFailureDate time.Time `gorm:"type:timestamp"`
	FailureProbability   float64   `gorm:"type:float"`
	CreatedAt            time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt            time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
	Equipment            Equipment `gorm:"foreignKey:EquipmentID"`
}

func (U *Prediction) BeforeCreate(tx *gorm.DB) (err error) {
	U.ID = uuid.New().String()
	return
}
