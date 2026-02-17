package usercontroller

import "github.com/labstack/echo/v4"

func (c *Controller) SetRoutesAdmin(e *echo.Echo) {
	userGroup := e.Group("/admin/users")

	userGroup.GET("/" ,c.GetAllUsers)
	userGroup.PATCH("/:userID" ,c.UpdateUserAdmin)
}