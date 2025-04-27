package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dzikuri/simple-withdraw-and-store-money/api/handler"
	"github.com/dzikuri/simple-withdraw-and-store-money/api/router"
	"github.com/dzikuri/simple-withdraw-and-store-money/config"
	"github.com/dzikuri/simple-withdraw-and-store-money/repository"
	"github.com/dzikuri/simple-withdraw-and-store-money/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

type APIserve struct {
	DB     *pgxpool.Pool
	Logger zerolog.Logger
}

func NewAPIServe(db *pgxpool.Pool, logger zerolog.Logger) *APIserve {
	return &APIserve{DB: db, Logger: logger}
}

func (r *APIserve) Serve() {
	// NOTE: Echo Framework
	e := echo.New()

	// NOTE: Use a zerolog middleware to log each request
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
	}))

	// NOTE: Health Check
	e.GET("/health", func(c echo.Context) error {
		r.Logger.Info().Msg("Health check called")
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "ok",
		})
	})

	// NOTE: Error Handling Page not found
	e.Any("/*", func(c echo.Context) error {
		r.Logger.Warn().Msg("Route not found")
		return c.JSON(http.StatusNotFound, map[string]interface{}{"remark": "Not found"})
	})

	// NOTE: Call Repository
	nasabahRepository := repository.NewNasabahRepository(r.DB, r.Logger)
	nasabahTransactionRepository := repository.NewNasabahTransactionRepository(r.DB, r.Logger)

	// NOTE: Call Service
	registerService := service.NewRegisterService(nasabahRepository, r.Logger)
	transactionService := service.NewTransactionService(nasabahTransactionRepository, nasabahRepository, r.Logger)

	// NOTE: Call Handler
	nasabahHandler := handler.NewNasabahHandler(registerService, transactionService, r.Logger)

	// NOTE: Call Router
	routeConfig := router.NewRouter(nasabahHandler)
	routeConfig.RegisterApiRouter(e)

	// NOTE: Handle graceful shutdown
	go func() {
		// Listen for interrupt or termination signals
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		sigReceived := <-sigs

		r.Logger.Info().Str("signal", sigReceived.String()).Msg("Received signal, starting graceful shutdown")

		// Attempt to stop the Echo server gracefully
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := e.Shutdown(ctx); err != nil {
			r.Logger.Error().Err(err).Msg("Error during server shutdown")

		}
		r.Logger.Info().Msg("Server gracefully stopped")

	}()

	// NOTE: Start the Echo server
	r.Logger.Info().Msg("Starting server")
	if err := e.Start(fmt.Sprintf(":%s", config.ConfigEnv.ApiPort)); err != nil {
		r.Logger.Fatal().Err(err).Msg("Error starting server")
	}
}
