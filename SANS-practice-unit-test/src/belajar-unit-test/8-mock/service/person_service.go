package service

import (
	"belajar-unit-test/belajar-unit-test/8-mock/entity"
	"belajar-unit-test/belajar-unit-test/8-mock/repository"
	"errors"
)

type PersonService struct {
	Repository repository.PersonRepository
}

func (service PersonService) Get(id string) (*entity.Person, error) {
	result := service.Repository.FindById(id)

	if result == nil {
		return nil, errors.New("NOT FOUND")
	} else {
		return result, nil
	}
}
