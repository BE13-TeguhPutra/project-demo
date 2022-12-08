package repository

import (
	"be13/project/features/feedback"
	_feedback "be13/project/features/feedback"

	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	UserID   uint
	Status   string
	MenteeID uint
	Comment  string
}

func CoretoModel(dataCore feedback.Core) Feedback {
	classGorm := Feedback{

		UserID:   dataCore.UserID,
		Status:   dataCore.Status,
		MenteeID: dataCore.MenteeID,
		Comment:  dataCore.Comment,
	}
	return classGorm
}

func (dataModel *Feedback) ModeltoCore() feedback.Core {
	return feedback.Core{
		ID:        dataModel.ID,
		UserID:    dataModel.UserID,
		Status:    dataModel.Status,
		MenteeID:  dataModel.MenteeID,
		Comment:   dataModel.Comment,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}

}

func toCoreList(models []Feedback) []feedback.Core {
	var feedCore []feedback.Core
	for _, v := range models {
		feedCore = append(feedCore, v.ModeltoCore())

	}
	return feedCore
}

func FromUserCore(dataCore _feedback.Core) Feedback { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	userGorm := Feedback{
		UserID:   dataCore.UserID,
		Status:   dataCore.Status,
		MenteeID: dataCore.MenteeID,
		Comment:  dataCore.Comment,
	} ///formating data berdasarkan data gorm dan kita mapping data yang kita butuhkan untuk inputan  klien
	return userGorm //insert user
}
