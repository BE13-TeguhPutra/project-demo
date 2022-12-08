package repository

import (
	"be13/project/features/feedback"
	"errors"

	"gorm.io/gorm"
)

type feedbackRepository struct {
	db *gorm.DB
}

func NewFeedback(db *gorm.DB) feedback.RepositoryInterface {
	return &feedbackRepository{
		db: db,
	}
}

// GetAllFeeds implements feedback.RepositoryInterface
func (repo *feedbackRepository) GetAllFeeds() (data []feedback.Core, err error) {
	var feeds []Feedback
	tx := repo.db.Find(&feeds)
	if tx.Error != nil {
		return nil, tx.Error
	}
	res := toCoreList(feeds) // model to core
	return res, nil
}

// GetFeedbyId implements feedback.RepositoryInterface
func (repo *feedbackRepository) GetFeedbyId(id int) (feedback.Core, error) {
	feeds := Feedback{}
	tx := repo.db.First(&feeds, id)
	if tx.Error != nil {
		return feedback.Core{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return feedback.Core{}, errors.New("id not found")

	}
	result := feeds.ModeltoCore()
	return result, nil
}

// AddFeedback implements feedback.RepositoryInterface
func (repo *feedbackRepository) AddFeedback(input feedback.Core) error {
	userGorm := FromUserCore(input) //dari gorm model ke user core yang ada di entities

	tx := repo.db.Create(&userGorm) // proses insert data

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// DeleteFeedback implements feedback.RepositoryInterface
func (repo *feedbackRepository) DeleteFeedback(id int) error {
	var Feeds Feedback
	tx := repo.db.Delete(&Feeds, id)
	if tx.Error != nil {

		return tx.Error
	}
	return nil
}

// // UpdateFeedback implements feedback.RepositoryInterface
func (repo *feedbackRepository) UpdateFeedback(id int, input feedback.Core) error {
	feedGorm := FromUserCore(input)

	tx := repo.db.Model(&feedGorm).Where("id = ?", id).Updates(&feedGorm)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
