package delivery

import (
	"be13/project/features/mentee"
	"be13/project/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MenteeDelivery struct {
	menteeServices mentee.ServiceInterface
}

func NewMentee(service mentee.ServiceInterface, e *echo.Echo) {
	handler := &MenteeDelivery{
		menteeServices: service,
	}

	e.POST("/mentees", handler.AddMentee)
	e.GET("/mentees", handler.GetAllmentee)
	e.GET("/mentee", handler.GetMentebyParam)
	e.GET("/mentee/:id/feedback", handler.GetMenteeFeedback)
	e.PUT("/mentees/:id", handler.UpdateMentee)
	e.DELETE("/mentees/:id", handler.DeleteMentee)

}

func (delivery *MenteeDelivery) GetAllmentee(c echo.Context) error {

	// name := c.QueryParam("name")
	result, err := delivery.menteeServices.GetAllmentee()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Data Users is empty"))
	}
	data := responseList(result)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success Get All mentee", data))

}

func (delivery *MenteeDelivery) GetMentebyParam(c echo.Context) error {
	// // idParam := c.Param("id")
	// id, errconv := strconv.Atoi(idParam)
	// if errconv != nil {
	// 	return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Convert"))
	// }
	name := c.QueryParam("name")
	// status := c.QueryParam("status")
	// category := c.QueryParam("category")
	// class := c.QueryParam("class")
	// classId, _ := strconv.Atoi(class)

	userId, err := delivery.menteeServices.GetMentebyParam(name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Id not Found"))
	}

	result := responseList(userId)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success Get mentee", result))

}

func (delivery *MenteeDelivery) AddMentee(c echo.Context) error {

	mentee := MenteeRequest{}
	errBind := c.Bind(&mentee)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error binding"))
	}

	result := mentee.reqToCore()

	err := delivery.menteeServices.AddMentee(result)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error insert into database"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Success Add Mantee"))

}

func (delivery *MenteeDelivery) UpdateMentee(c echo.Context) error {
	mentee := MenteeRequest{}
	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Convert"))
	}

	errBind := c.Bind(&mentee)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error binding"))
	}

	result := mentee.reqToCore()
	err := delivery.menteeServices.UpdateMentee(id, result)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Update"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Success Update"))

}
func (delivery *MenteeDelivery) DeleteMentee(c echo.Context) error {
	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Convert"))
	}
	data, err := delivery.menteeServices.DeleteMentee(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Delete Mentee"))
	}

	result := coreToResponse(data)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success Delete Mentee", result))

}

func (delivery *MenteeDelivery) GetMenteeFeedback(c echo.Context) error {
	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Convert"))
	}
	// name := c.QueryParam("name")
	// status := c.QueryParam("status")
	// category := c.QueryParam("category")
	// class := c.QueryParam("class")
	// classId, _ := strconv.Atoi(class)

	userId, err := delivery.menteeServices.GetMenteeFeedback(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Id not Found"))
	}

	result := coreToResponse(userId)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success Get mentee", result))

}
