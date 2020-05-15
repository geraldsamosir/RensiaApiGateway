package router

import (
	"net/http"

	"github.com/geraldsamosir/RensiaApiGateway/AuthServices/delivery/http/utils"
	"github.com/labstack/echo"
)

type (
	RbacServiceRouter struct {
	}

	RbacService interface {
		CheckAllowAccess(ctx echo.Context) error
	}
)

func NewRbacServiceRouter() *RbacServiceRouter {
	return &RbacServiceRouter{}
}

func (rbacf *RbacServiceRouter) CheckAllowAccess(ctx echo.Context) error {
	data := "authServices"
	return utils.Response(http.StatusOK, data, nil, ctx)

}
