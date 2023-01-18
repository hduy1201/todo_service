package apis

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"itss.edu.vn/todo_service/business"
	"itss.edu.vn/todo_service/core"
	"itss.edu.vn/todo_service/models"
)

func NewAuthApi(endpoint string, server *core.Server) *echo.Group {
	apis := server.Echo.Group(endpoint)
	authBusiness := business.NewAuthBusiness(server)

	apis.POST("/register", func(c echo.Context) error {
		var regUser models.UserRegistrationRequest
		if err := c.Bind(&regUser); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if err := c.Validate(&regUser); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if err := authBusiness.Register(&regUser); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, regUser)
	})

	return apis
}
