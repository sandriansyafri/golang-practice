package repository

import (
	"belajar-unit-test/belajar-unit-test/8-mock/entity"

	"github.com/stretchr/testify/mock"
)

type PersonRepositoryMock struct {
	Mock mock.Mock
}

func (mock PersonRepositoryMock) FindById(id string) *entity.Person {
	args := mock.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	} else {
		result := args.Get(0).(entity.Person)
		return &result
	}
}
