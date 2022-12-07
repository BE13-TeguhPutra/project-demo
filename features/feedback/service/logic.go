package service

import (
	"be13/project/features/feedback"

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

// AddFeedback implements feedback.ServiceInterface
func (*feedbackService) AddFeedback(input feedback.Core) error {
	panic("unimplemented")
}

// DeleteFeedback implements feedback.ServiceInterface
func (*feedbackService) DeleteFeedback(id int) (feedback.Core, error) {
	panic("unimplemented")
}

// UpdateFeedback implements feedback.ServiceInterface
func (*feedbackService) UpdateFeedback(id int, input feedback.Core) error {
	panic("unimplemented")
}
