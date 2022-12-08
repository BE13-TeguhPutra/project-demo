package delivery

import (
	"be13/project/features/feedback"
	"time"
)

type FeedbackResponse struct {
	ID        uint   `json:"feedback_id"`
	UserID    uint   `json:"user_id"`
	Status    string `json:"status"`
	MenteeID  uint   `json:"mentee_id"`
	Comment   string `json:"comment"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CoreToResponse(core feedback.Core) FeedbackResponse {
	response := FeedbackResponse{
		ID:        core.ID,
		UserID:    core.UserID,
		Status:    core.Status,
		MenteeID:  core.MenteeID,
		Comment:   core.Comment,
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
	}
	return response

}

func responseList(listRes []feedback.Core) []FeedbackResponse {
	var feedList []FeedbackResponse
	for _, v := range listRes {
		feedList = append(feedList, CoreToResponse(v))

	}
	return feedList

}
