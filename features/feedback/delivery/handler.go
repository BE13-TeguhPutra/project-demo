package delivery

import (
	"be13/project/features/feedback"
	"be13/project/helper"
	"be13/project/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FeedbackDelivery struct {
	feedbackServices feedback.ServiceInterface
}

func NewFeedback(service feedback.ServiceInterface, e *echo.Echo) {
	handler := &FeedbackDelivery{
		feedbackServices: service,
	}

	e.POST("/feedback", handler.Addfeedback, middlewares.JWTMiddleware())
	e.GET("/feedback", handler.GetAllFeeds, middlewares.JWTMiddleware())
	e.GET("/feedback/:id", handler.GetFeedbyId, middlewares.JWTMiddleware())
	e.PUT("/feedback", handler.Updatefeedback, middlewares.JWTMiddleware())
	e.DELETE("/feedback/:id", handler.Deletefeedback, middlewares.JWTMiddleware())

}

func (delivery *FeedbackDelivery) Addfeedback(c echo.Context) error {

	// roletoken := middlewares.ExtractTokenUserRole(c)
	// }

	InputFeedback := FeedbackRequest{}
	errbind := c.Bind(&InputFeedback)

	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("erorr read data"+errbind.Error()))
	}
	dataCore := FeedbackRequestToUserCore(InputFeedback) //data mapping yang diminta create
	errResultCore := delivery.feedbackServices.AddFeedback(dataCore)
	if errResultCore != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("erorr read data"+errResultCore.Error()))
	}
	return c.JSON(http.StatusCreated, helper.PesanSuksesHelper("berhasil create user"))
}
func (delivery *FeedbackDelivery) Updatefeedback(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	userInput := FeedbackRequest{}
	errBind := c.Bind(&userInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("Error binding data "+errBind.Error()))
	}

	dataCore := FeedbackRequestToUserCore(userInput)
	err := delivery.feedbackServices.UpdateFeedback(id, dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.PesanGagalHelper("Failed update data"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.PesanSuksesHelper("success Update data"))
}

func (delivery *FeedbackDelivery) Deletefeedback(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result := delivery.feedbackServices.DeleteFeedback(id) //memanggil fungsi service yang ada di folder service

	if result != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("erorr read data"))
	}

	return c.JSON(http.StatusCreated, helper.PesanSuksesHelper("success Hapus data"))
}

func (delivery *FeedbackDelivery) GetAllFeeds(c echo.Context) error {
	result, err := delivery.feedbackServices.GetAllFeeds()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("Data feedback is empty"))
	}
	data := responseList(result)

	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("Success Get All Feedbacks", data))

}

func (delivery *FeedbackDelivery) GetFeedbyId(c echo.Context) error {
	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("error Convert"+errconv.Error()))
	}
	// name := c.QueryParam("name")

	userId, err := delivery.feedbackServices.GetFeedbyId(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("fedback id not Found"+err.Error()))
	}

	result := CoreToResponse(userId)
	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("Success Get Feedback", result))

}
