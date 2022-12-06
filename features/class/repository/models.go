package repository

import (
	"be13/project/features/class"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	ClassName string
	UserID    uint
}

func CoretoModel(dataCore class.Core) Class {
	classGorm := Class{

		ClassName: dataCore.ClassName,
		UserID:    dataCore.UserID,
	}
	return classGorm
}

func (dataModel *Class) ModeltoCore() class.Core {
	return class.Core{
		ID:        dataModel.ID,
		UserID:    dataModel.UserID,
		ClassName: dataModel.ClassName,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}

}

func toCoreList(models []Class) []class.Core {
	var userCore []class.Core
	for _, v := range models {
		userCore = append(userCore, v.ModeltoCore())

	}
	return userCore
}
