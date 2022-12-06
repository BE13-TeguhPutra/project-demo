package service

import (
	"be13/project/features/user"
	"errors"

	"github.com/go-playground/validator/v10"
)

// bisnis Logic
type UserService struct {
	userRepository user.RepositoryEntities //data repository dri entities
	validate       *validator.Validate
}

// Login implements user.ServiceEntities

// Login implements user.ServiceEntities

func New(repo user.RepositoryEntities) user.ServiceEntities { //dengan kembalian user.service
	return &UserService{
		userRepository: repo,
		validate:       validator.New(),
	}
}

// Create implements user.ServiceEntities
func (service *UserService) Create(input user.CoreUser) (err error) {

	if validateERR := service.validate.Struct(input); validateERR != nil {
		return validateERR
	}

	_, errCreate := service.userRepository.Create(input)
	if errCreate != nil {
		return errors.New("GAGAL MENAMBAH DATA , QUERY ERROR")
	}

	return nil
}

// func (service *UserService) Login(input user.CoreUser) (token string, err error) {
// 	if errValidate := service.validate.Struct(input); errValidate != nil {
// 		log.Error(errValidate.Error())
// 		return "", errors.New("Failed to Login. Error validate input. Please check your input.")
// 	}

// 	result, errLogin := service.userRepository.FindUser(input.Email)
// 	if errLogin != nil {
// 		log.Error(errLogin.Error())
// 		if strings.Contains(errLogin.Error(), "table") {
// 			return "", errors.New("Failed to Login. Error on request. Please contact your administrator.")
// 		} else if strings.Contains(errLogin.Error(), "found") {
// 			return "", errors.New("Failed to Login. Email not found. Please check password again.")
// 		} else {
// 			return "", errors.New("Failed to Login. Other Error. Please contact your administrator.")
// 		}
// 	}

// 	errCheckPass := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(input.Password))
// 	// fmt.Println("Data Core = ", dataCore, "\n\n\n")
// 	// fmt.Println("Result = ", result, "\n\n\n")
// 	if errCheckPass != nil {
// 		log.Error(errCheckPass.Error())
// 		return "", errors.New("Failed to Login. Password didn't match. Please check password again.")
// 	}

// 	token, errToken := middlewares.CreateToken(int(result.ID), result.Role)
// 	if errToken != nil {
// 		log.Error(errToken.Error())
// 		return "", errors.New("Failed to login. Error on generate token. Please check password again.")
// 	}

// 	return token, nil
// }

// DeleteById implements user.ServiceEntities
// func (*UserService) DeleteById(id int) error {
// 	panic("unimplemented")
// }

// // GetAll implements user.ServiceEntities
// func (*UserService) GetAll() (data []user.CoreUser, err error) {
// 	panic("unimplemented")
// }

// // GetById implements user.ServiceEntities
// func (*UserService) GetById(id int) (data user.CoreUser, err error) {
// 	panic("unimplemented")
// }

// // Update implements user.ServiceEntities
// func (*UserService) Update(id int, input user.CoreUser) error {
// 	panic("unimplemented")
// }
