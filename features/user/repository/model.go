package repository

import (
	"be13/project/features/class/repository"
	_user "be13/project/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(500)"`
	Email    string
	Phone    string
	Address  string
	Status   string
	Role     string
	Classes  []repository.Class
	// Feedbacks []_feedback.Feedback
}

func FromUserCore(dataCore _user.CoreUser) User { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	userGorm := User{
		FullName: dataCore.FullName,
		Email:    dataCore.Email, //mapping data core ke data gorm model
		Password: dataCore.Password,
		Phone:    dataCore.Phone,
		Address:  dataCore.Address,
		Status:   dataCore.Status,
		Role:     dataCore.Role,
	} ///formating data berdasarkan data gorm dan kita mapping data yang kita butuhkan untuk inputan  klien
	return userGorm //insert user
}
func (dataModel *User) ModelsToCore() _user.CoreUser { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return _user.CoreUser{

		FullName: dataModel.FullName,
		Email:    dataModel.Email, //mapping data core ke data gorm model
		Password: dataModel.Password,
		Phone:    dataModel.Phone,
		Address:  dataModel.Address,
		Status:   dataModel.Status,
		Role:     dataModel.Role,
	}
}
func ListModelTOCore(dataModel []User) []_user.CoreUser { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []_user.CoreUser
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}
