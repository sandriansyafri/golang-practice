package service

import (
	"app/entity"
	"app/repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("NOT FOUND")
	} else {
		return category, nil
	}
}
