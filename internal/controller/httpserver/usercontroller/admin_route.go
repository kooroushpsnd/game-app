package usercontroller

import "github.com/labstack/echo/v4"

func (c *Controller) SetRoutesAdmin(e *echo.Echo) {
	userGroup := e.Group("/admin/users")

	userGroup.GET("" ,c.GetAllUsers)
	// userGroup.PUT("/:userID" ,c.UpdateUser)
}