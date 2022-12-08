package feedback

import "time"

type Core struct {
	ID        uint
	Status    string `validate:"required"`
	UserID    uint   `validate:"required"`
	MenteeID  uint
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceInterface interface {
	GetAllFeeds() (data []Core, err error)
	AddFeedback(input Core) error
	GetFeedbyId(id int) (Core, error)
	DeleteFeedback(id int) error
	UpdateFeedback(id int, input Core) error
}

type RepositoryInterface interface {
	GetAllFeeds() (data []Core, err error)
	AddFeedback(input Core) error
	GetFeedbyId(id int) (Core, error)
	DeleteFeedback(id int) error
	UpdateFeedback(id int, input Core) error
}
