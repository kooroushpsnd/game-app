package usercontroller

import (
	userdto "goProject/internal/dto/user"
	"goProject/internal/pkg/helper"
	"goProject/internal/pkg/httpmsg"
	"goProject/internal/pkg/mapper"
	"goProject/internal/validator"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (controller *Controller) GetAllUsers(c echo.Context) error{
	req, err := helper.BindValidateQuery[userdto.GetAllRequestUserDto](c)
	if err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  validator.FieldErrors(err),
		})
	}

	resp, err := controller.userSvc.GetAllUsers(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (controller *Controller) UpdateUserAdmin(c echo.Context) error{
	idStr := c.Param("userID")
    userID, err := strconv.Atoi(idStr)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest)
    }

	var req userdto.UpdateRequestAdminDto
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

	resp, err := controller.userSvc.Update(c.Request().Context(),uint(userID) ,mapper.AdminDtoToPatch(req))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
