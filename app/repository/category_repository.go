package repository

import (
	"jokibro/app/models"

	"gorm.io/gorm"
)

type categoryRepository struct {
	DB *gorm.DB
}

type Icategory interface {
	FindAll(search string) (res []models.Category, err error)
}

func NewCategory(DB *gorm.DB) Icategory {
	return &categoryRepository{DB: DB}
}

func (repo categoryRepository) FindAll(search string) (res []models.Category, err error) {
	data := repo.DB.Where("name LIKE ?", "%"+search+"%").Find(&res)
	return res, data.Error
}
