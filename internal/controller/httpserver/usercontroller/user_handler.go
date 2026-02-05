package usercontroller

import (
	userdto "goProject/internal/dto/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (controller Controller) userLogin(c echo.Context) error {
	var req userdto.UserLoginRequestDto
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	

	resp, err := controller.userSvc.Login(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
