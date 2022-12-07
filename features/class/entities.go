package class

import "time"

type Core struct {
	ID        uint
	ClassName string `validate:"required"`
	UserID    uint   `validate:"required"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceInterface interface {
	GetAllClass() (data []Core, err error)
	AddClass(input Core) error
	GetClassbyId(name string) (Core, error)
	DeleteClass(id int) (Core, error)
	UpdateClass(id int, input Core) error
}

type RepositoryInterface interface {
	GetAllClass() (data []Core, err error)
	AddClass(input Core) error
	GetClassbyId(name string) (Core, error)
	DeleteClass(id int) (Core, error)
	UpdateClass(id int, input Core) error
}
