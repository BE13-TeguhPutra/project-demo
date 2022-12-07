package mentee

import (
	"be13/project/features/feedback"
	"time"
)

type Core struct {
	ID                uint
	Name              string `validate:"required"`
	Address           string `validate:"required"`
	HomeAddress       string `validate:"required"`
	Email             string `validate:"required"`
	Gender            string `validate:"required"`
	Telegram          string `validate:"required"`
	Phone             string `validate:"required"`
	ManteeStatus      string `validate:"required"`
	EmergencyPhone    string `validate:"required"`
	EmergencyName     string `validate:"required"`
	EmergencyRelation string `validate:"required"`
	Category          string `validate:"required"`
	Major             string `validate:"required"`
	Graduate          string `validate:"required"`
	ClassID           uint   `validate:"required"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Feedbacks         []feedback.Core
}

type ServiceInterface interface {
	GetAllmentee() (data []Core, err error)
	AddMentee(input Core) error
	GetMentebyParam(name string) ([]Core, error)
	DeleteMentee(id int) (Core, error)
	UpdateMentee(id int, input Core) error
	GetMenteeFeedback(id uint) (Core, error)
}

type RepositoryInterface interface {
	GetAllmentee() (data []Core, err error)
	AddMentee(input Core) error
	GetMentebyParam(name string) ([]Core, error)
	DeleteMentee(id int) (Core, error)
	UpdateMentee(id int, input Core) error
	GetMenteeFeedback(id uint) (Core, error)
}
