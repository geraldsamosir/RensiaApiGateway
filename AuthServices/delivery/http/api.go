package http

import (
	"net/http"

	"github.com/geraldsamosir/RensiaApiGateway/AuthServices/delivery/http/router"
	"github.com/geraldsamosir/RensiaApiGateway/AuthServices/delivery/http/utils"
	"github.com/labstack/echo"
)

type HttpServer interface {
	RunServer(port string)
	router(r *echo.Echo) *echo.Echo
}

func handler(ctx echo.Context) error {
	data := "authServices"
	return utils.Response(http.StatusOK, data, nil, ctx)
}

func routers(r *echo.Echo) *echo.Echo {
	//handler
	authHandler := router.NewAuthRouter()
	rbacFrontendHandler := router.NewRbacFrontendRouter()
	rbacServicesHandler := router.NewRbacServiceRouter()

	//root router
	roohandler := r.Group("/auth")

	//routers
	roohandler.POST("/", authHandler.LoginUser)
	roohandler.POST("/register", authHandler.RegisterUser)
	roohandler.POST("/soc-med/login", authHandler.LoginUserSocMed)
	roohandler.POST("/soc-med", authHandler.RegisterUserSocmed)
	roohandler.GET("/rbac-services", rbacServicesHandler.CheckAllowAccess)
	roohandler.GET("/rbac-frontend", rbacFrontendHandler.CheckAllowAccess)

	//return router
	return r
}

func RunServer(port string) {
	r := echo.New()
	routers(r)
	r.Start(":" + port)
}
