package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/dzikuri/simple-withdraw-and-store-money/model"
	"github.com/dzikuri/simple-withdraw-and-store-money/service"
	"github.com/dzikuri/simple-withdraw-and-store-money/util"
	"github.com/go-playground/validator/v10"
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

	param := model.GetSaldoParameter{
		RekeningNumber: c.Param("no_rekening"),
	}

	// Validate the struct
	if err := util.Validator.Struct(param); err != nil {

		// Check if the error is a validation error
		if _, ok := err.(*validator.ValidationErrors); !ok {

			c.Logger().Info(err)

			// Custom error handling
			return c.JSON(http.StatusBadRequest, map[string]string{
				"remark": "Nomor rekening tidak valid",
			})
		}

		// Custom error handling
		return c.JSON(http.StatusBadRequest, map[string]string{
			"remark": err.Error(),
		})
	}

	result, err := h.TransactionService.CheckSaldo(c.Request().Context(), param.RekeningNumber)
	if err != nil {
		// Check if error is "no rows in result set"
		if errors.Is(err, sql.ErrNoRows) {

			c.Logger().Info(err.Error())

			return c.JSON(http.StatusBadRequest, map[string]string{
				"remark": "Nomor rekening tidak dikenali",
			})
		}

		c.Logger().Warn(err.Error())

		// Other unknown/internal error
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"remark": "Terjadi kesalahan internal",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"saldo": result})

}

func (h *NasabahHandler) Deposit(c echo.Context) error {
	var req model.TransactionPayload

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"remark": err.Error()})
	}

	// Validate the struct
	if err := util.Validator.Struct(req); err != nil {

		// Check if the error is a validation error
		if _, ok := err.(*validator.ValidationErrors); !ok {

			c.Logger().Info(err)

			// Custom error handling
			return c.JSON(http.StatusBadRequest, map[string]string{
				"remark": "Nomor rekening tidak valid",
			})
		}

		// Custom error handling
		return c.JSON(http.StatusBadRequest, map[string]string{
			"remark": err.Error(),
		})
	}

	amount, err := h.TransactionService.DepositMoney(c.Request().Context(), req.NasabahId, req.Amount)
	if err != nil {
		// Check if error is "no rows in result set"
		if errors.Is(err, sql.ErrNoRows) {

			c.Logger().Info(err.Error())

			return c.JSON(http.StatusBadRequest, map[string]string{
				"remark": "Nomor rekening tidak dikenali",
			})
		}

		c.Logger().Warn(err.Error())

		// Other unknown/internal error
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"remark": "Terjadi kesalahan internal",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"saldo": amount})
}
func (h *NasabahHandler) Withdraw(c echo.Context) error {
	var req model.TransactionPayload

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"remark": err.Error()})
	}

	// Validate the struct
	if err := util.Validator.Struct(req); err != nil {

		// Check if the error is a validation error
		if _, ok := err.(*validator.ValidationErrors); !ok {

			c.Logger().Info(err)

			// Custom error handling
			return c.JSON(http.StatusBadRequest, map[string]string{
				"remark": "Nomor rekening tidak valid",
			})
		}

		// Custom error handling
		return c.JSON(http.StatusBadRequest, map[string]string{
			"remark": err.Error(),
		})
	}

	amount, err := h.TransactionService.WithdrawMoney(c.Request().Context(), req.NasabahId, req.Amount)
	if err != nil {
		// Check if error is "no rows in result set"
		if errors.Is(err, sql.ErrNoRows) {

			c.Logger().Info(err.Error())

			return c.JSON(http.StatusBadRequest, map[string]string{
				"remark": "Nomor rekening tidak dikenali",
			})
		}

		if errors.Is(err, util.ErrInsufficientBalance) {

			c.Logger().Info(err.Error())

			return c.JSON(http.StatusBadRequest, map[string]string{
				"remark": "Saldo tidak mencukupi",
			})

		}

		c.Logger().Warn(err.Error())

		// Other unknown/internal error
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"remark": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"saldo": amount})
}
