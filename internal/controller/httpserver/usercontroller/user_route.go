package usercontroller

import (
	"goProject/internal/controller/httpserver/middleware"

	"github.com/labstack/echo/v4"
)

func (c *Controller) SetRoutesUser(e *echo.Echo) {
	userGroup := e.Group("/users")

	userGroup.GET("/profile", c.Profile, middleware.Auth(c.authSvc, c.authConfig))
	userGroup.POST("/login", c.Login)
	userGroup.POST("/signup", c.Signup)
}
