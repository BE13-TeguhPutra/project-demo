package factory

import (
	deliveryAuth "be13/project/features/auth/delivery"
	repoAuth "be13/project/features/auth/repository"
	serviceAuth "be13/project/features/auth/service"
	deliveryClass "be13/project/features/class/delivery"
	repoClass "be13/project/features/class/repository"
	serviceClass "be13/project/features/class/service"
	userDelivery "be13/project/features/user/delivery"
	userRepo "be13/project/features/user/repository" //alias
	userService "be13/project/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userRepofaktory := userRepo.New(db) //menginiasialisasi func new yang ada di repository
	userServiceFaktory := userService.New(userRepofaktory)
	userDelivery.New(userServiceFaktory, e)

	classRepoFactory := repoClass.NewClass(db)
	classServiceFactory := serviceClass.NewClass(classRepoFactory)
	deliveryClass.NewClass(classServiceFactory, e)

	authRepoFactory := repoAuth.NewAuth(db)
	authServiceFactory := serviceAuth.NewAuth(authRepoFactory)
	deliveryAuth.NewAuth(authServiceFactory, e)

}
