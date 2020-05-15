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
	roothandler := r.Group("/auth")

	//routers
	roothandler.POST("/", authHandler.LoginUser)
	roothandler.POST("/register", authHandler.RegisterUser)
	roothandler.POST("/soc-med/login", authHandler.LoginUserSocMed)
	roothandler.POST("/soc-med", authHandler.RegisterUserSocmed)
	roothandler.GET("/rbac-services", rbacServicesHandler.CheckAllowAccess)
	roothandler.GET("/rbac-frontend", rbacFrontendHandler.CheckAllowAccess)

	//return router
	return r
}

func RunServer(port string) {
	r := echo.New()
	routers(r)
	r.Start(":" + port)
}
