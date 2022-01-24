package service

import (
	"belajar-unit-test/belajar-unit-test/8-mock/entity"
	"belajar-unit-test/belajar-unit-test/8-mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var personRepository = &repository.PersonRepositoryMock{Mock: mock.Mock{}} // mock
var personService = PersonService{Repository: personRepository}

func TestPersonNotFound(t *testing.T) {
	personRepository.Mock.On("FindById", "1").Return(nil) // set mock

	result, err := personService.Get("2")
	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func TestPersonFound(t *testing.T) {
	person := entity.Person{
		Name: "A",
		Id:   "1",
	}

	personRepository.Mock.On("FindById", "2").Return(person) // set mock

	result, err := personService.Get("3")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, person.Id, result.Id)
	assert.Equal(t, person.Name, result.Name)
}
