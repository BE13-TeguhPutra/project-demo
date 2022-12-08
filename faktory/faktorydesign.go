package factory

import (
	deliveryAuth "be13/project/features/auth/delivery"
	repoAuth "be13/project/features/auth/repository"
	serviceAuth "be13/project/features/auth/service"
	deliveryClass "be13/project/features/class/delivery"
	repoClass "be13/project/features/class/repository"
	serviceClass "be13/project/features/class/service"
	deliveryFeedback "be13/project/features/feedback/delivery"
	repoFeedback "be13/project/features/feedback/repository"
	serviceFeedback "be13/project/features/feedback/service"
	deliveryMentee "be13/project/features/mentee/delivery"
	repoMentee "be13/project/features/mentee/repository"
	serviceMentee "be13/project/features/mentee/service"
	userDelivery "be13/project/features/user/delivery"
	userRepo "be13/project/features/user/repository"
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

	menteeRepoFactory := repoMentee.NewMentee(db)
	menteeServiceFactory := serviceMentee.NewMentee(menteeRepoFactory)
	deliveryMentee.NewMentee(menteeServiceFactory, e)

	feedbackRepoFactory := repoFeedback.NewFeedback(db)
	feedbackServiceFactory := serviceFeedback.NewFeedback(feedbackRepoFactory)
	deliveryFeedback.NewFeedback(feedbackServiceFactory, e)

}
