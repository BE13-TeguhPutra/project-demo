package delivery

///handler = controller
import (
	"be13/project/features/user"
	"be13/project/helper"
	"be13/project/middlewares"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserDeliv struct {
	UserService user.ServiceEntities
}

func New(Service user.ServiceEntities, e *echo.Echo) {
	handler := &UserDeliv{
		UserService: Service,
	}
	// e.GET("/user", handler.GetAll) // memanggil func getall
	e.POST("/user", handler.Create, middlewares.JWTMiddleware())
	// e.POST("/auth", handler.Login)
	// e.PUT("/user/:id", handler.Update)
	// e.GET("/user/:id", handler.GetById)
	// e.DELETE("/user/:id", handler.DeleteById)
}
func (delivery *UserDeliv) Create(c echo.Context) error {

	roletoken := middlewares.ExtractTokenUserRole(c)
	log.Println("Role Token", roletoken)
	if roletoken != "admin" {
		return c.JSON(http.StatusUnauthorized, helper.PesanGagalHelper("tidak bisa diakses khusus admin!!!"))
	}

	Inputuser := UserRequest{} //penangkapan data user reques dari entities user
	errbind := c.Bind(&Inputuser)

	generatePass := user.Bcript(Inputuser.Password)
	Inputuser.Password = generatePass

	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("erorr read data"+errbind.Error()))
	}
	dataCore := UserRequestToUserCore(Inputuser) //data mapping yang diminta create
	errResultCore := delivery.UserService.Create(dataCore)
	if errResultCore != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("erorr read data"+errResultCore.Error()))
	}
	return c.JSON(http.StatusCreated, helper.PesanSuksesHelper("berhasil create user"))
}

// func (handler *UserDeliv) Login(c echo.Context) error {
// 	userInput := AuthRequest{}
// 	errBind := c.Bind(&userInput)
// 	if errBind != nil {
// 		// return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 		// 	"message": "failed to get bind data",
// 		// })
// 		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("erorr read data"))
// 	}

// 	dataCore := ToCore(userInput)
// 	token, err := handler.UserService.Login(dataCore)

// 	if err != nil {
// 		// return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 		// 	"message": "failed to get token data" + err.Error(),
// 		// })
// 		return c.JSON(http.StatusInternalServerError, helper.PesanGagalHelper("Failed to Login. "+err.Error()))
// 	}
// return c.JSON(http.StatusOK, map[string]interface{}{
// 	"message": "success",
// 	"name":    name,
// 	"token":   result,
// })
// 	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("login success", map[string]interface{}{
// 		"token": token,
// 	}))
// }
