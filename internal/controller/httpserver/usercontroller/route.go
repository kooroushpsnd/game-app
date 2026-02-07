package usercontroller

import "github.com/labstack/echo/v4"

func (c Controller) SetRoutes(e *echo.Echo){
	userGroup := e.Group("/users")

	// userGroup.GET("/profile")
	userGroup.POST("/login" ,c.Login)
	userGroup.POST("/signup" ,c.Signup)
}