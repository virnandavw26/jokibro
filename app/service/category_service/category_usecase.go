package service

import (
	"jokibro/app/http/viewmodel"
	"jokibro/app/repository"

	"gorm.io/gorm"
)

type CategoryUsecase struct {
	DB *gorm.DB
}

func (uc CategoryUsecase) FindAll(search string) (data []viewmodel.CategoryVM, err error) {
	repo := repository.NewCategory(uc.DB)
	dataModel, err := repo.FindAll(search)
	if err != nil {
		return data, err
	}

	for _, d := range dataModel {
		data = append(data, viewmodel.CategoryVM{
			ID:        d.ID,
			Name:      d.Name,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
			DeletedAt: d.DeletedAt,
		})
	}

	return data, err
}
