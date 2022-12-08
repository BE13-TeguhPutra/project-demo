package service

import (
	"be13/project/features/feedback"
	"errors"

	"github.com/go-playground/validator/v10"
)

type feedbackService struct {
	feedbackRepository feedback.RepositoryInterface
	validate           *validator.Validate
}

func NewFeedback(repo feedback.RepositoryInterface) feedback.ServiceInterface {
	return &feedbackService{
		feedbackRepository: repo,
		validate:           validator.New(),
	}
}

// GetAllFeeds implements feedback.ServiceInterface
func (service *feedbackService) GetAllFeeds() (data []feedback.Core, err error) {
	data, err = service.feedbackRepository.GetAllFeeds()
	return data, err
}

// GetFeedbyId implements feedback.ServiceInterface
func (service *feedbackService) GetFeedbyId(id int) (feedback.Core, error) {
	data, errData := service.feedbackRepository.GetFeedbyId(id)
	if errData != nil {
		return feedback.Core{}, errors.New("error get id")
	}
	return data, nil
}

// AddFeedback implements feedback.ServiceInterface
func (service *feedbackService) AddFeedback(input feedback.Core) error {
	errCreate := service.feedbackRepository.AddFeedback(input)
	if errCreate != nil {
		return errCreate
	}

	return nil
}

// DeleteFeedback implements feedback.ServiceInterface
func (service *feedbackService) DeleteFeedback(id int) error {
	data := service.feedbackRepository.DeleteFeedback(id) // memanggil struct entities repository yang ada di entities yang berisi coding logic
	return data
}

// UpdateFeedback implements feedback.ServiceInterface
func (service *feedbackService) UpdateFeedback(id int, input feedback.Core) error {
	errUpdate := service.feedbackRepository.UpdateFeedback(id, input)
	if errUpdate != nil {
		return errors.New("GAGAL mengupdate data , QUERY ERROR")
	}

	return nil
}
