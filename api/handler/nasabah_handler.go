package handler

import (
	"net/http"

	"github.com/dzikuri/simple-withdraw-and-store-money/model"
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
	var req model.CreateNasabah

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"remark": err.Error()})
	}

	rekeningNumber, err := h.RegisterService.RegisterNasabah(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"remark": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"rekening_number": rekeningNumber})
}

func (h *NasabahHandler) GetSaldo(c echo.Context) error {
	return nil
}

func (h *NasabahHandler) Withdraw(c echo.Context) error {
	return nil
}

func (h *NasabahHandler) Deposit(c echo.Context) error {
	return nil
}
