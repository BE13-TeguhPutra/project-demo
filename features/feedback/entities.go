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
	// GetAllClass() (data []Core, err error)
	AddFeedback(input Core) error
	// GetClassbyId(name string) (Core, error)
	DeleteFeedback(id int) (Core, error)
	UpdateFeedback(id int, input Core) error
}

type RepositoryInterface interface {
	// GetAllClass() (data []Core, err error)
	AddFeedback(input Core) error
	// GetClassbyId(name string) (Core, error)
	DeleteFeedback(id int) (Core, error)
	UpdateFeedback(id int, input Core) error
}
