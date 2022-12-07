package delivery

import "be13/project/features/user"

type UserRespon struct {
	// ID    uint   `json:"id`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     string `json:"role" form:"role"`
	Address  string `json:"address"`
	Status   string `json:"status"`
	// Book  []BookRespon `json:"book"`
}

func UserCoreToUserRespon(dataCore user.CoreUser) UserRespon { // data user core yang ada di controller yang memanggil user repository
	return UserRespon{
		// ID:    dataCore.ID,
		FullName: dataCore.FullName,
		Email:    dataCore.Email, //mapping data core ke data gorm model
		Phone:    dataCore.Phone,
		Address:  dataCore.Address,
		Status:   dataCore.Status,
		Role:     dataCore.Role,
	}
}
func ListUserCoreToUserRespon(dataCore []user.CoreUser) []UserRespon { //data user.core data yang diambil dari entities ke respon struct
	var ResponData []UserRespon

	for _, value := range dataCore { //memanggil paramete data core yang berisi data user core
		ResponData = append(ResponData, UserCoreToUserRespon(value)) // mengambil data mapping dari user core to respon
	}
	return ResponData
}
