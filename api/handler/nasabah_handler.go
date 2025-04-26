package handler

import (
	"net/http"

	"github.com/dzikuri/simple-withdraw-and-store-money/service"
	"github.com/labstack/echo/v4"
)

type NasabahHandler struct {
	RegisterService    service.RegisterService
	TransactionService service.TransactionService
}

func NewNasabahHandler(registerService service.RegisterService, transactionService service.TransactionService) *NasabahHandler {
	return &NasabahHandler{RegisterService: registerService, TransactionService: transactionService}
}

func (h *NasabahHandler) CreateNasabah(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Nasabah created!",
	})
}

func (h *NasabahHandler) GetSaldo(c echo.Context) {

}

func (h *NasabahHandler) Withdraw(c echo.Context) {

}

func (h *NasabahHandler) Deposit(c echo.Context) {

}
