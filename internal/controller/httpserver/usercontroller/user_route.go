package usercontroller

import (
	"goProject/internal/controller/httpserver/middleware"

	"github.com/labstack/echo/v4"
)

func (c *Controller) SetRoutesUser(e *echo.Echo) {
	userGroup := e.Group("/users")

	userGroup.POST("/login", c.Login)
	userGroup.POST("/signup", c.Signup)
	userGroup.POST("/send_email_code" ,c.SendEmailCode)
	userGroup.POST("/verify_email" ,c.VerifyEmail)

	userGroup.GET("/profile", c.Profile, middleware.Auth(c.authSvc, c.authConfig))
	
	userGroup.PATCH("/" ,c.UpdateUser, middleware.Auth(c.authSvc, c.authConfig))
}
