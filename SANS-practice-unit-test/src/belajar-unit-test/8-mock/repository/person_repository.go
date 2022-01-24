package repository

import "belajar-unit-test/belajar-unit-test/8-mock/entity"

type PersonRepository interface {
	FindById(id string) *entity.Person
}
