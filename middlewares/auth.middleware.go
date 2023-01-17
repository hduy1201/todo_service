package middlewares

import (
	"github.com/labstack/echo/v4"
	"itss.edu.vn/todo_service/utils"
)

func NewAuthMiddleware() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			rawToken := c.Request().Header.Get("Authorization")
			if rawToken == "" {
				return c.JSON(401, "Unauthorized")
			}
			uid, err := utils.JWTVerify(rawToken)
			if err != nil {
				return c.JSON(401, "Unauthorized")
			}
			c.Set("uid", uid)
			return next(c)
		}
	}
}
