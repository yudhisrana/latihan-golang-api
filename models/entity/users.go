package entity

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"type:varchar(100);not null"`
	Email     string `json:"email" gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}