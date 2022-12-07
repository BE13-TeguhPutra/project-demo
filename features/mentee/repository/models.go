package repository

import (
	"be13/project/features/feedback"
	"be13/project/features/feedback/repository"
	"be13/project/features/mentee"

	"gorm.io/gorm"
)

type Mentee struct {
	gorm.Model
	Name              string
	Address           string
	HomeAddress       string
	Email             string
	Gender            string
	Telegram          string
	Phone             string
	ManteeStatus      string
	EmergencyPhone    string
	EmergencyName     string
	EmergencyRelation string
	Category          string
	Major             string
	Graduate          string
	ClassID           uint
	Feedbacks         []repository.Feedback
}

func CoretoModel(dataCore mentee.Core) Mentee {
	classGorm := Mentee{

		Name:              dataCore.Name,
		Address:           dataCore.Address,
		HomeAddress:       dataCore.HomeAddress,
		Email:             dataCore.Email,
		Gender:            dataCore.Gender,
		Telegram:          dataCore.Telegram,
		Phone:             dataCore.Phone,
		ManteeStatus:      dataCore.ManteeStatus,
		EmergencyPhone:    dataCore.EmergencyPhone,
		EmergencyName:     dataCore.EmergencyName,
		EmergencyRelation: dataCore.EmergencyRelation,
		Category:          dataCore.Category,
		Major:             dataCore.Major,
		Graduate:          dataCore.Graduate,
		ClassID:           dataCore.ClassID,
	}
	return classGorm
}

func (dataModel *Mentee) ModeltoCore() mentee.Core {
	return mentee.Core{
		ID:                dataModel.ID,
		Name:              dataModel.Name,
		Address:           dataModel.Address,
		HomeAddress:       dataModel.HomeAddress,
		Email:             dataModel.Email,
		Gender:            dataModel.Gender,
		Telegram:          dataModel.Telegram,
		Phone:             dataModel.Phone,
		ManteeStatus:      dataModel.ManteeStatus,
		EmergencyPhone:    dataModel.EmergencyPhone,
		EmergencyName:     dataModel.EmergencyName,
		EmergencyRelation: dataModel.EmergencyRelation,
		Category:          dataModel.Category,
		Major:             dataModel.Major,
		Graduate:          dataModel.Graduate,
		ClassID:           dataModel.ClassID,
		CreatedAt:         dataModel.CreatedAt,
		UpdatedAt:         dataModel.UpdatedAt,
		Feedbacks:         LoadFeedsModeltoCore(dataModel.Feedbacks),
	}

}

func toCoreList(models []Mentee) []mentee.Core {
	var menteeCore []mentee.Core
	for _, v := range models {
		menteeCore = append(menteeCore, v.ModeltoCore())

	}
	return menteeCore
}

func LoadFeedsModeltoCore(model []repository.Feedback) []feedback.Core {
	var core []feedback.Core
	for _, v := range model {
		core = append(core, v.ModeltoCore())
	}
	return core

}
