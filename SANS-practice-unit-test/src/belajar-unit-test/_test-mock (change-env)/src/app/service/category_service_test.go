package service

import (
	"app/entity"
	"app/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryServiceNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")

	assert.Nil(t, category)
	assert.NotNil(t, err)

}

func TestCategoryServiceFound(t *testing.T) {
	category := entity.Category{
		Id:   "2",
		Name: "Category-1",
	}
	category2 := entity.Category{
		Id:   "2",
		Name: "Category-3",
	}

	categoryRepository.Mock.On("FindById", "2").Return(category)
	categoryRepository.Mock.On("FindById", "3").Return(category2)

	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
