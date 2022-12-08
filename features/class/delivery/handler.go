package delivery

import (
	"be13/project/features/class"
	"be13/project/middlewares"
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

	e.POST("/classes", handler.AddClass, middlewares.JWTMiddleware())
	e.GET("/classes", handler.GetAllClass, middlewares.JWTMiddleware())
	e.GET("/classes/:id", handler.GetClassbyName, middlewares.JWTMiddleware())
	e.PUT("/classes/:id", handler.UpdateClass, middlewares.JWTMiddleware())
	e.DELETE("/classes/:id", handler.DeleteClass, middlewares.JWTMiddleware())

}

func (delivery *ClassDelivery) AddClass(c echo.Context) error {
	// roletoken := middlewares.ExtractTokenUserRole(c)
	// log.Println("Role Token", roletoken)
	// if roletoken != "admin" {
	// 	return c.JSON(http.StatusUnauthorized, helper.FailedResponse("tidak bisa diakses khusus admin!!!"))
	// }
	class := ClassRequest{}
	errBind := c.Bind(&class)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error binding"+errBind.Error()))
	}

	result := class.reqToCore()

	err := delivery.classServices.AddClass(result)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error insert into database"+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Success Add Class"))

}

func (delivery *ClassDelivery) GetAllClass(c echo.Context) error {
	result, err := delivery.classServices.GetAllClass()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Data Users is empty"+err.Error()))
	}
	data := responseList(result)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success Get All Users", data))

}

func (delivery *ClassDelivery) GetClassbyName(c echo.Context) error {
	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Convert"+errconv.Error()))
	}
	// name := c.QueryParam("name")

	userId, err := delivery.classServices.GetClassbyId(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Id not Found"+err.Error()))
	}

	result := coreToResponse(userId)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success Get Class", result))

}

func (delivery *ClassDelivery) UpdateClass(c echo.Context) error {
	// roletoken := middlewares.ExtractTokenUserRole(c)
	// log.Println("Role Token", roletoken)
	// if roletoken != "admin" {
	// 	return c.JSON(http.StatusUnauthorized, helper.FailedResponse("tidak bisa diakses khusus admin!!!"))
	// }
	class := ClassRequest{}
	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Convert"+errconv.Error()))
	}

	errBind := c.Bind(&class)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error binding"+errBind.Error()))
	}

	result := class.reqToCore()
	err := delivery.classServices.UpdateClass(id, result)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Update"+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Success Update"))

}

func (delivery *ClassDelivery) DeleteClass(c echo.Context) error {
	// roletoken := middlewares.ExtractTokenUserRole(c)
	// log.Println("Role Token", roletoken)
	// if roletoken != "admin" {
	// 	return c.JSON(http.StatusUnauthorized, helper.FailedResponse("tidak bisa diakses khusus admin!!!"))
	// }
	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Convert"+errconv.Error()))
	}
	data, err := delivery.classServices.DeleteClass(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Delete"+err.Error()))
	}

	result := coreToResponse(data)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success Delete Class", result))

}
