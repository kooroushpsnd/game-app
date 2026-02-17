package usercontroller

import (
	userdto "goProject/internal/dto/user"
	"goProject/internal/pkg/claim"
	"goProject/internal/pkg/httpmsg"
	"goProject/internal/pkg/mapper"
	"goProject/internal/validator"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (controller *Controller) Login(c echo.Context) error {
	var req userdto.LoginRequestDto
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

	resp, err := controller.userSvc.Login(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (controller *Controller) Signup(c echo.Context) error {
	var req userdto.SignupRequestDto
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

	resp, err := controller.userSvc.Register(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (controller *Controller) Profile(c echo.Context) error{
	claims := claim.GetClaimsFromEchoContext(c)

	resp ,err := controller.userSvc.GetProfile(
		c.Request().Context() ,
		userdto.GetProfileRequestDto{UserID: claims.UserID},
	)
	if err != nil {
		msg ,code := httpmsg.Error(err)
		return echo.NewHTTPError(code ,msg)
	}

	return c.JSON(http.StatusOK ,resp)
}

func (controller *Controller) UpdateUser(c echo.Context) error{
	claims := claim.GetClaimsFromEchoContext(c)

	var req userdto.UpdateRequestUserDto
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

	resp, err := controller.userSvc.Update(c.Request().Context(),claims.UserID ,mapper.UserDtoToPatch(req))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (controller *Controller) VerifyEmail(c echo.Context) error{
	
}