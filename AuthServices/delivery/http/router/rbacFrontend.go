package router

import (
	"net/http"

	"github.com/geraldsamosir/RensiaApiGateway/AuthServices/delivery/http/utils"
	"github.com/labstack/echo"
)

type RbacFrontendRouter struct {
}

type RbacFrontend interface {
	CheckAllowAccess(ctx echo.Context) error
}

func NewRbacFrontendRouter() *RbacFrontendRouter {
	return &RbacFrontendRouter{}
}

func (rbacf *RbacFrontendRouter) CheckAllowAccess(ctx echo.Context) error {
	data := "rbac frontend"
	return utils.Response(http.StatusOK, data, nil, ctx)

}
