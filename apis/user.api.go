package apis

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"itss.edu.vn/todo_service/business"
	"itss.edu.vn/todo_service/core"
	"itss.edu.vn/todo_service/middlewares"
	"itss.edu.vn/todo_service/models"
)

func NewUserApi(endpoint string, server *core.Server) *echo.Group {
	apis := server.Echo.Group(endpoint)
	userBusiness := business.NewUserBusiness(server)

	apis.Use(middlewares.NewAuthMiddleware())

	apis.POST("/", func(c echo.Context) error {
		user := &models.User{}
		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if err := c.Validate(user); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if err := userBusiness.Create(user); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, user)
	})

	apis.GET("/", func(c echo.Context) error {
		id := c.QueryParam("id")
		username := c.QueryParam("username")

		user, err := userBusiness.Get(id, username)
		if err != nil {
			return c.JSON(http.StatusNotFound, err)
		}
		return c.JSON(http.StatusOK, user)
	})

	apis.PUT("/", func(c echo.Context) error {
		user := &models.User{}
		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if err := c.Validate(user); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if err := userBusiness.Update(user); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, user)
	})

	apis.DELETE("/", func(c echo.Context) error {
		id := c.QueryParam("id")
		if id == "" {
			return c.JSON(http.StatusBadRequest, "Id is required")
		}
		if err := userBusiness.Delete(id); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.NoContent(http.StatusOK)
	})
	return apis
}
