package http

import (
	"net/http"

	"github.com/labstack/echo"
)

type HttpServer interface {
	RunServer(port string)
	router()
}

type res struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"error"`
}

func handler(ctx echo.Context) error {
	data := "authServices"
	payload := res{Status: http.StatusOK, Data: data, Error: ""}
	return ctx.JSON(http.StatusOK, payload)
}

func router(r *echo.Echo) *echo.Echo {
	r.GET("/", handler)
	r.POST("/auth", handler)
	r.POST("/Auth/soc-med", handler)
	r.GET("/rbac-services", handler)
	r.GET("/rbac-frontend", handler)
	return r
}

func RunServer(port string) {
	r := echo.New()
	router(r)
	r.Start(":" + port)
}
