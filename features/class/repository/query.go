package repository

import (
	"be13/project/features/class"
	"errors"

	"gorm.io/gorm"
)

type classRepository struct {
	db *gorm.DB
}

func NewClass(db *gorm.DB) class.RepositoryInterface {
	return &classRepository{
		db: db,
	}
}

// AddClass implements class.RepositoryInterface
func (repo *classRepository) AddClass(input class.Core) error {
	classGorm := CoretoModel(input)
	tx := repo.db.Create(&classGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

// DeleteClass implements class.RepositoryInterface
func (repo *classRepository) DeleteClass(id int) (class.Core, error) {
	classes := Class{}
	tx := repo.db.Delete(&classes, id)
	if tx.Error != nil {
		return class.Core{}, tx.Error
	}

	tx1 := repo.db.Unscoped().Where("id=?", id).Find(&classes)
	if tx1.Error != nil {
		return class.Core{}, tx1.Error
	}
	if tx.RowsAffected == 0 {
		return class.Core{}, errors.New("id not found")

	}
	result := classes.ModeltoCore()
	return result, nil
}

// GetAllClass implements class.RepositoryInterface
func (repo *classRepository) GetAllClass() (data []class.Core, err error) {
	var class []Class
	tx := repo.db.Find(&class)
	if tx.Error != nil {
		return nil, tx.Error
	}
	res := toCoreList(class) // model to core
	return res, nil
}

// GetClassbyId implements class.RepositoryInterface
func (repo *classRepository) GetClassbyId(id uint) (class.Core, error) {
	classes := Class{}
	tx := repo.db.First(&classes, id)
	if tx.Error != nil {
		return class.Core{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return class.Core{}, errors.New("id not found")

	}
	result := classes.ModeltoCore()
	return result, nil
}

// UpdateClass implements class.RepositoryInterface
func (repo *classRepository) UpdateClass(id int, input class.Core) error {
	class := CoretoModel(input)
	tx := repo.db.Model(class).Where("id = ?", id).Updates(class)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}
