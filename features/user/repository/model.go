package repository

import (
	"be13/project/features/class/repository"
	"be13/project/features/user"
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

func (data User) toCore() user.CoreUser {
	return user.CoreUser{

		FullName:  data.FullName,
		Email:     data.Email,
		Password:  data.Password,
		Role:      data.Role,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
