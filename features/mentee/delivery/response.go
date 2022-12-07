package delivery

import (
	"be13/project/features/feedback"
	"be13/project/features/feedback/delivery"
	"be13/project/features/mentee"
)

type MenteeResponse struct {
	ID                uint   `json:"id_mantee"`
	Name              string `json:"name"`
	Address           string `json:"address"`
	HomeAddress       string `json:"home_address"`
	Email             string `json:"email"`
	Gender            string `json:"gender"`
	Telegram          string `json:"telegram"`
	Phone             string `json:"phone"`
	ManteeStatus      string `json:"mentee_status"`
	EmergencyPhone    string `json:"emergency_phone"`
	EmergencyName     string `json:"emergency_name"`
	EmergencyRelation string `json:"emergency_relation"`
	Category          string `json:"category"`
	Major             string `json:"major"`
	Graduate          string `json:"graduate"`
	ClassID           uint   `json:"class_id"`
	// Feedbacks         []delivery.FeedbackResponse
}

func coreToResponse(core mentee.Core) MenteeResponse {
	response := MenteeResponse{
		ID:                core.ID,
		Name:              core.Name,
		Address:           core.Address,
		HomeAddress:       core.HomeAddress,
		Email:             core.Email,
		Gender:            core.Gender,
		Telegram:          core.Telegram,
		Phone:             core.Phone,
		ManteeStatus:      core.ManteeStatus,
		EmergencyPhone:    core.EmergencyPhone,
		EmergencyName:     core.EmergencyName,
		EmergencyRelation: core.EmergencyRelation,
		Category:          core.Category,
		Major:             core.Major,
		Graduate:          core.Graduate,
		ClassID:           core.ClassID,
		// Feedbacks:         LoadFeedsCoretoResponse(core.Feedbacks),
	}
	return response

}

func responseList(listRes []mentee.Core) []MenteeResponse {
	var classList []MenteeResponse
	for _, v := range listRes {
		classList = append(classList, coreToResponse(v))

	}
	return classList

}

func LoadFeedsCoretoResponse(core []feedback.Core) []delivery.FeedbackResponse {
	var response []delivery.FeedbackResponse
	for _, v := range core {
		response = append(response, delivery.CoreToResponse(v))
	}
	return response

}
