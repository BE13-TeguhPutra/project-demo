package delivery

import (
	"be13/project/features/class"
	"be13/project/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ClassDelivery struct {
	classServices class.ServiceInterface
}

func NewClass(service class.ServiceInterface, e *echo.Echo) {
	handler := &ClassDelivery{
		classServices: service,
	}

	e.POST("/classes", handler.AddClass)
	e.GET("/classes", handler.GetAllClass)
	e.GET("/classes/", handler.GetClassbyName)
	e.PUT("/classes/:id", handler.UpdateClass)
	e.DELETE("/classes/:id", handler.DeleteClass)

}

func (delivery *ClassDelivery) AddClass(c echo.Context) error {
	// role := middlewares.ExtractTokenUserRole(c)
	// // fmt.Println(role)
	// // if role != "super admin" {
	// // 	return c.JSON(http.StatusUnauthorized, helper.FailedResponse("Failed role is not super admin"))

	// // }
	class := ClassRequest{}
	errBind := c.Bind(&class)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error binding"))
	}

	result := class.reqToCore()

	err := delivery.classServices.AddClass(result)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error insert into database"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Success Add Class"))

}

func (delivery *ClassDelivery) GetAllClass(c echo.Context) error {
	result, err := delivery.classServices.GetAllClass()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Data Users is empty"))
	}
	data := responseList(result)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success Get All Users", data))

}

func (delivery *ClassDelivery) GetClassbyName(c echo.Context) error {
	// idParam := c.Param("id")
	// id, errconv := strconv.Atoi(idParam)
	// if errconv != nil {
	// 	return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Convert"))
	// }
	name := c.QueryParam("name")

	userId, err := delivery.classServices.GetClassbyId(name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Id not Found"))
	}

	result := coreToResponse(userId)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success Get cLASS", result))

}

func (delivery *ClassDelivery) UpdateClass(c echo.Context) error {
	class := ClassRequest{}
	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Convert"))
	}

	errBind := c.Bind(&class)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error binding"))
	}

	result := class.reqToCore()
	err := delivery.classServices.UpdateClass(id, result)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Update"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Success Update"))

}

func (delivery *ClassDelivery) DeleteClass(c echo.Context) error {
	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Convert"))
	}
	data, err := delivery.classServices.DeleteClass(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Delete"))
	}

	result := coreToResponse(data)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success Delete Class", result))

}
