package router

import (
	"net/http"

	"github.com/geraldsamosir/RensiaApiGateway/AuthServices/delivery/http/utils"
	"github.com/labstack/echo"
)

type AuthRouter struct {
}

type auth interface {
	RegisterUser(ctx echo.Context) error
	LoginUser(ctx echo.Context) error
	ForgotPassword(ctc echo.Context) error
	LoginUserSocMed(ctc echo.Context) error
	RegisterUserSocmed(ctc echo.Context) error
}

func NewAuthRouter() *AuthRouter {
	return &AuthRouter{}
}

func (a *AuthRouter) RegisterUser(ctx echo.Context) error {
	data := "authServices register"
	return utils.Response(http.StatusOK, data, nil, ctx)
}

func (a *AuthRouter) LoginUserSocMed(ctx echo.Context) error {
	data := "authServices login with socmed"
	return utils.Response(http.StatusOK, data, nil, ctx)
}

func (a *AuthRouter) RegisterUserSocmed(ctx echo.Context) error {
	data := "authServices login with socmed"
	return utils.Response(http.StatusOK, data, nil, ctx)
}

func (a *AuthRouter) LoginUser(ctx echo.Context) error {
	data := "authServices login"
	return utils.Response(http.StatusOK, data, nil, ctx)
}

func (a *AuthRouter) ForgotPassword(ctx echo.Context) error {
	data := "authServices forgot"
	return utils.Response(http.StatusOK, data, nil, ctx)
}
