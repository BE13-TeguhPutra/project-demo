package delivery

import (
	"be13/project/features/user"
)

type UserRequest struct {
	FullName string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Role     string `json:"role" form:"role"`
}

func UserRequestToUserCore(data UserRequest) user.CoreUser {
	return user.CoreUser{
		FullName: data.FullName,
		Password: data.Password,
		Email:    data.Email,
		Phone:    data.Phone,
		Address:  data.Address,
		Role:     data.Role,
	}
}
