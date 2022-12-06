package delivery

import "be13/project/features/class"

type ClassResponse struct {
	ClassName string `json:"class_name"`
	UserID    uint   `json:"user_id"`
}

func coreToResponse(core class.Core) ClassResponse {
	response := ClassResponse{
		ClassName: core.ClassName,
		UserID:    core.UserID,
	}
	return response

}

func responseList(listRes []class.Core) []ClassResponse {
	var classList []ClassResponse
	for _, v := range listRes {
		classList = append(classList, coreToResponse(v))

	}
	return classList

}
