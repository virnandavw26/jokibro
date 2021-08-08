package viewmodel

import (
	"time"

	"gorm.io/gorm"
)

type CategoryVM struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
