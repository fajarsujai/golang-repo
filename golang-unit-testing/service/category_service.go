package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"errors"
)

type CategoryService struct {
	repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.repository.FindById(id)

	if category == nil {
		return nil, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}