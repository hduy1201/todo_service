package apis

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"itss.edu.vn/todo_service/core"
	"itss.edu.vn/todo_service/models"
)

func NewTaskApi(endpoint string, server *core.Server) *echo.Group {
	apis := server.Echo.Group(endpoint)

	tasks := make([]*models.Task, 0)

	apis.POST("/", func(c echo.Context) error {
		task := &models.Task{}

		if err := c.Bind(task); err != nil {
			return echo.ErrBadRequest //truyen body k hop le
		}

		if err := c.Validate(task); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}

		if task.Id == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Id is required",
			})
		}

		tasks = append(tasks, task)
		return c.NoContent(http.StatusCreated)
	})

	apis.GET("/all", func(c echo.Context) error {
		return c.JSON(http.StatusOK, tasks)
	})

	return apis
}
