package delivery

import "be13/project/features/user"

type AuthRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToCore(userReq AuthRequest) user.CoreUser {
	userCore := user.CoreUser{
		Email:    userReq.Email,
		Password: userReq.Password,
	}
	return userCore
}
