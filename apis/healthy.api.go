package apis

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"itss.edu.vn/todo_service/core"
	"itss.edu.vn/todo_service/middlewares"
)

func NewHealthyAPI(endpoint string, server *core.Server) *echo.Group {
	apis := server.Echo.Group(endpoint)

	apis.Use(middlewares.NewAuthMiddleware())

	apis.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "OK",
		})
	})

	return apis
}
