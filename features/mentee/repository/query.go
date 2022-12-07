package repository

import (
	"be13/project/features/mentee"
	"errors"

	"gorm.io/gorm"
)

type menteeRepository struct {
	db *gorm.DB
}

func NewMentee(db *gorm.DB) mentee.RepositoryInterface {
	return &menteeRepository{
		db: db,
	}
}

// GetMentebFeedback implements mentee.RepositoryInterface
func (repo *menteeRepository) GetMenteeFeedback(id uint) (mentee.Core, error) {
	mentees := Mentee{}

	tx := repo.db.First(&mentees, id)
	if tx.Error != nil {
		return mentee.Core{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return mentee.Core{}, errors.New("id not found")

	}
	result := mentees.ModeltoCore()
	return result, nil
}

// AddMentee implements mentee.RepositoryInterface
func (repo *menteeRepository) AddMentee(input mentee.Core) error {
	menteeGorm := CoretoModel(input)
	tx := repo.db.Create(&menteeGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

// DeleteMentee implements mentee.RepositoryInterface
func (repo *menteeRepository) DeleteMentee(id int) (mentee.Core, error) {
	mentees := Mentee{}
	tx := repo.db.Delete(&mentees, id)
	if tx.Error != nil {
		return mentee.Core{}, tx.Error
	}

	tx1 := repo.db.Unscoped().Where("id=?", id).Find(&mentees)
	if tx1.Error != nil {
		return mentee.Core{}, tx1.Error
	}
	if tx.RowsAffected == 0 {
		return mentee.Core{}, errors.New("id not found")

	}
	result := mentees.ModeltoCore()
	return result, nil
}

// GetAllmentee implements mentee.RepositoryInterface
func (repo *menteeRepository) GetAllmentee() (data []mentee.Core, err error) {
	var mentee []Mentee
	tx := repo.db.Find(&mentee)
	if tx.Error != nil {
		return nil, tx.Error
	}
	res := toCoreList(mentee) // model to core
	return res, nil
}

// GetMentebyParam implements mentee.RepositoryInterface
func (repo *menteeRepository) GetMentebyParam(name string) ([]mentee.Core, error) {
	var mentees []Mentee
	tx := repo.db.Where("name=?", name).Find(&mentees)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("id not found")

	}
	result := toCoreList(mentees)
	return result, nil
}

// UpdateMentee implements mentee.RepositoryInterface
func (repo *menteeRepository) UpdateMentee(id int, input mentee.Core) error {
	mantee := CoretoModel(input)
	tx := repo.db.Model(mantee).Where("id = ?", id).Updates(mantee)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}
