package delivery

import (
	"be13/project/features/mentee"
)

type MenteeRequest struct {
	Name              string `json:"name" form:"name"`
	Address           string `json:"address" form:"address"`
	HomeAddress       string `json:"home_address" form:"home_address"`
	Email             string `json:"email" form:"email"`
	Gender            string `json:"gender" form:"gender"`
	Telegram          string `json:"telegram" form:"telegram"`
	Phone             string `json:"phone" form:"phone"`
	ManteeStatus      string `json:"mentee_status" form:"mentee_status"`
	EmergencyPhone    string `json:"emergency_phone" form:"emergency_phone"`
	EmergencyName     string `json:"emergency_name" form:"emergency_name"`
	EmergencyRelation string `json:"emergency_relation" form:"emergency_relation"`
	Category          string `json:"category" form:"category"`
	Major             string `json:"major" form:"major"`
	Graduate          string `json:"graduate" form:"graduate"`
	ClassID           uint   `json:"class_id" form:"class_id"`
}

func (data *MenteeRequest) reqToCore() mentee.Core {
	return mentee.Core{
		Name:              data.Name,
		Address:           data.Address,
		HomeAddress:       data.HomeAddress,
		Email:             data.Email,
		Gender:            data.Gender,
		Telegram:          data.Telegram,
		Phone:             data.Phone,
		ManteeStatus:      data.ManteeStatus,
		EmergencyPhone:    data.EmergencyPhone,
		EmergencyName:     data.EmergencyName,
		EmergencyRelation: data.EmergencyRelation,
		Category:          data.Category,
		Major:             data.Major,
		Graduate:          data.Graduate,
		ClassID:           data.ClassID,
	}
}
