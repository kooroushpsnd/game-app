package usercontroller

import (
	userdto "goProject/internal/dto/user"
	"goProject/internal/pkg/httpmsg"
	"goProject/internal/validator"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (controller *Controller) GetAllUsers(c echo.Context) error{
	var req userdto.GetAllRequestUserDto
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if err := c.Validate(&req); err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg ,
			"errors": validator.FieldErrors(err),
		})
	}

	resp, err := controller.userSvc.GetAllUsers(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

// func (c *Controller) UpdateUser(e echo.Context) error{}
