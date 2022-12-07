package repository

import (
	"be13/project/features/user"
	"errors"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// GetAll implements user.RepositoryEntities

// FindUser implements user.RepositoryEntities

// FindUser implements user.RepositoryEntities

func New(db *gorm.DB) user.RepositoryEntities { // user.repository mengimplementasikan interface repository yang ada di entities
	return &userRepository{
		db: db,
	}

}

// Create implements user.RepositoryEntities
func (repo *userRepository) Create(input user.CoreUser) (row int, err error) {
	userGorm := FromUserCore(input) //dari gorm model ke user core yang ada di entities

	tx := repo.db.Create(&userGorm) // proses insert data

	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("Insert failed")
	}
	return int(tx.RowsAffected), nil
}
func (repo *userRepository) GetAll() (data []user.CoreUser, err error) {
	var users []User //mengambil data gorm model(model.go)
	tx := repo.db.Unscoped().Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var DataCore = ListModelTOCore(users) //mengambil data dari gorm model(file repository(model.go))

	return DataCore, nil

}

// DeleteById implements user.RepositoryEntities
// func (*userRepository) DeleteById(id int) error {
// 	panic("unimplemented")
// }

// // GetAll implements user.RepositoryEntities
// func (*userRepository) GetAll() (data []user.CoreUser, err error) {
// 	panic("unimplemented")
// }

// // GetById implements user.RepositoryEntities
// func (*userRepository) GetById(id int) (data user.CoreUser, err error) {
// 	panic("unimplemented")
// }

// // Update implements user.RepositoryEntities
// func (*userRepository) Update(id int, input user.CoreUser) error {
// 	panic("unimplemented")
// }
