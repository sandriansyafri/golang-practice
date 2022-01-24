package repository

import "app/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
