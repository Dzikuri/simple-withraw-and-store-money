package router

import (
	"github.com/dzikuri/simple-withdraw-and-store-money/api/handler"
	"github.com/labstack/echo/v4"
)

type RouteConfig struct {
	NasabahHandler *handler.NasabahHandler
}

func NewRouter(nasabahHandler *handler.NasabahHandler) *RouteConfig {
	return &RouteConfig{
		NasabahHandler: nasabahHandler,
	}
}

func (r *RouteConfig) RegisterApiRouter(e *echo.Echo) {
	e.POST("/daftar", r.NasabahHandler.CreateNasabah)
	e.GET("/saldo/:no_rekening", r.NasabahHandler.GetSaldo)
}
