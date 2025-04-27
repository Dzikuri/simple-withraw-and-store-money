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
	"github.com/rs/zerolog"
)

type NasabahHandler struct {
	RegisterService    service.RegisterService
	TransactionService service.TransactionService
	Logger             zerolog.Logger
}

func NewNasabahHandler(registerService service.RegisterService, transactionService service.TransactionService, logger zerolog.Logger) *NasabahHandler {
	return &NasabahHandler{RegisterService: registerService, TransactionService: transactionService, Logger: logger}
}

func (h *NasabahHandler) CreateNasabah(c echo.Context) error {
	var req model.CreateNasabah

	if err := c.Bind(&req); err != nil {
		h.Logger.Warn().Err(err).Msg("Handler: Error binding request")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"remark": "Request tidak valid"})
	}

	// NOTE: Validate the struct
	if err := util.Validator.Struct(req); err != nil {
		if _, ok := err.(*validator.ValidationErrors); !ok {
			h.Logger.Warn().Err(err).Msg("Handler: Error binding request")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"remark": "Request tidak valid"})
		}

		h.Logger.Warn().Err(err).Msg("Handler: Error binding request")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"remark": "Request tidak valid"})
	}

	rekeningNumber, err := h.RegisterService.RegisterNasabah(c.Request().Context(), &req)
	if err != nil {

		if errors.Is(err, util.ErrNasabahAlreadyExist) {
			h.Logger.Warn().Err(err).Msg("Handler: Error binding request")

			return c.JSON(http.StatusConflict, map[string]interface{}{"remark": "NIK atau nomor handphone sudah terdaftar"})
		}

		h.Logger.Warn().Err(err).Msg("Handler: Internal server error")

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"remark": "Terjadi kesalahan pada server"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"rekening_number": rekeningNumber})
}

func (h *NasabahHandler) GetSaldo(c echo.Context) error {

	param := model.GetSaldoParameter{
		RekeningNumber: c.Param("no_rekening"),
	}

	// NOTE: Validate the struct
	if err := util.Validator.Struct(param); err != nil {

		// Check if the error is a validation error
		if _, ok := err.(*validator.ValidationErrors); !ok {

			h.Logger.Warn().Err(err).Msg("Handler: Error binding request")

			// Custom error handling
			return c.JSON(http.StatusBadRequest, map[string]string{
				"remark": "Nomor rekening tidak valid",
			})
		}

		h.Logger.Warn().Err(err).Msg("Handler: Error binding request")

		// Custom error handling
		return c.JSON(http.StatusBadRequest, map[string]string{
			"remark": "Nomor rekening tidak valid",
		})
	}

	result, err := h.TransactionService.CheckSaldo(c.Request().Context(), param.RekeningNumber)
	if err != nil {
		// Check if error is "no rows in result set"
		if errors.Is(err, sql.ErrNoRows) {

			h.Logger.Warn().Err(err).Msg("Handler: Error binding request")

			return c.JSON(http.StatusBadRequest, map[string]string{
				"remark": "Nomor rekening tidak dikenali",
			})
		}

		h.Logger.Warn().Err(err).Msg("Handler: Error binding request")

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
		h.Logger.Warn().Err(err).Msg("Handler: Error binding request")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"remark": "Request tidak valid"})
	}

	// NOTE: Validate the struct
	if err := util.Validator.Struct(req); err != nil {

		// Check if the error is a validation error
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "NasabahId":
					h.Logger.Warn().Err(err).Msg("Handler: Error binding request")

					return c.JSON(http.StatusBadRequest, map[string]string{
						"remark": "Nomor rekening tidak valid",
					})
				case "Amount":
					h.Logger.Warn().Err(err).Msg("Handler: Error binding request")

					if fieldError.Tag() == "gt" {
						return c.JSON(http.StatusBadRequest, map[string]string{
							"remark": "Nominal harus lebih besar dari 0",
						})
					}

					return c.JSON(http.StatusBadRequest, map[string]string{
						"remark": "Nominal tidak valid",
					})
				}
			}
		}

		// Fallback unknown error
		h.Logger.Warn().Err(err).Msg("Handler: Error binding request")
		return c.JSON(http.StatusBadRequest, map[string]string{
			"remark": "Request tidak valid",
		})
	}

	amount, err := h.TransactionService.DepositMoney(c.Request().Context(), req.NasabahId, req.Amount)
	if err != nil {
		// Check if error is "no rows in result set"
		if errors.Is(err, sql.ErrNoRows) {

			h.Logger.Warn().Err(err).Msg("Handler: Error binding request")

			return c.JSON(http.StatusBadRequest, map[string]string{
				"remark": "Nomor rekening tidak dikenali",
			})
		}

		h.Logger.Warn().Err(err).Msg("Handler: Error binding request")

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
		h.Logger.Warn().Err(err).Msg("Handler: Error binding request")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"remark": "Request tidak valid"})
	}

	// NOTE: Validate the struct
	if err := util.Validator.Struct(req); err != nil {

		// Check if the error is a validation error
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "NasabahId":
					h.Logger.Warn().Err(err).Msg("Handler: Error binding request")

					return c.JSON(http.StatusBadRequest, map[string]string{
						"remark": "Nomor rekening tidak valid",
					})
				case "Amount":
					h.Logger.Warn().Err(err).Msg("Handler: Error binding request")

					if fieldError.Tag() == "gt" {
						return c.JSON(http.StatusBadRequest, map[string]string{
							"remark": "Nominal harus lebih besar dari 0",
						})
					}
					return c.JSON(http.StatusBadRequest, map[string]string{
						"remark": "Nominal tidak valid",
					})
				}
			}
		}

		// Fallback unknown error
		h.Logger.Warn().Err(err).Msg("Handler: Error binding request")

		return c.JSON(http.StatusBadRequest, map[string]string{
			"remark": "Request tidak valid",
		})
	}

	amount, err := h.TransactionService.WithdrawMoney(c.Request().Context(), req.NasabahId, req.Amount)
	if err != nil {
		// Check if error is "no rows in result set"
		if errors.Is(err, sql.ErrNoRows) {

			h.Logger.Warn().Err(err).Msg("Handler: Error binding request")

			return c.JSON(http.StatusBadRequest, map[string]string{
				"remark": "Nomor rekening tidak dikenali",
			})
		}

		if errors.Is(err, util.ErrInsufficientBalance) {

			h.Logger.Warn().Err(err).Msg("Handler: Error binding request")

			return c.JSON(http.StatusBadRequest, map[string]string{
				"remark": "Saldo tidak mencukupi",
			})

		}

		h.Logger.Warn().Err(err).Msg("Handler: Internal Server Error")

		// Other unknown/internal error
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"remark": "Terjadi kesalahan internal",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"saldo": amount})
}
