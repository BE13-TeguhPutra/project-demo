package service

import (
	"be13/project/features/class"
	"errors"

	"github.com/go-playground/validator/v10"
)

type classService struct {
	classRepository class.RepositoryInterface
	validate        *validator.Validate
}

func NewClass(repo class.RepositoryInterface) class.ServiceInterface {
	return &classService{
		classRepository: repo,
		validate:        validator.New(),
	}
}

// AddClass implements class.ServiceInterface
func (service *classService) AddClass(input class.Core) error {
	err := service.classRepository.AddClass(input)
	if err != nil {
		return errors.New("failed to add class")
	}
	return nil
}

// DeleteClass implements class.ServiceInterface
func (service *classService) DeleteClass(id int) (class.Core, error) {
	data, err := service.classRepository.DeleteClass(id)
	return data, err
}

// GetAllClass implements class.ServiceInterface
func (service *classService) GetAllClass() (data []class.Core, err error) {
	data, err = service.classRepository.GetAllClass()
	return data, err
}

// GetClassbyId implements class.ServiceInterface
func (service *classService) GetClassbyId(id uint) (class.Core, error) {
	data, errData := service.classRepository.GetClassbyId(id)
	if errData != nil {
		return class.Core{}, errors.New("error get id")
	}
	return data, nil
}

// UpdateClass implements class.ServiceInterface
func (service *classService) UpdateClass(id int, input class.Core) error {
	err := service.classRepository.UpdateClass(id, input)
	if err != nil {
		return errors.New("id not found")
	}
	return nil
}
