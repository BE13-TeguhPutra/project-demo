package delivery

import "be13/project/features/class"

type ClassRequest struct {
	ClassName string `json:"class_name" form:"class_name"`
	UserID    uint   `json:"user_id" form:"user_id"`
}

func (req *ClassRequest) reqToCore() class.Core {
	return class.Core{
		// ClassName: req.ClassName,
		ClassName: req.ClassName,
		UserID:    req.UserID,
	}

}
