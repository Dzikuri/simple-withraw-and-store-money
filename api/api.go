package api

import (
	"context"
	"fmt"
	"log"
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
)

type APIserve struct {
	DB *pgxpool.Pool
}

func NewAPIServe(db *pgxpool.Pool) *APIserve {
	return &APIserve{DB: db}
}

func (r *APIserve) Serve() {
	// NOTE: Echo Framework
	e := echo.New()

	e.Use(middleware.Logger())

	// NOTE: Health Check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "ok",
		})
	})

	// NOTE: Error Handling Page not found
	e.Any("/*", func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"remark": "Not found"})
	})

	// NOTE: Call Repository
	nasabahRepository := repository.NewNasabahRepository(r.DB)
	nasabahTransactionRepository := repository.NewNasabahTransactionRepository(r.DB)

	// NOTE: Call Service
	registerService := service.NewRegisterService(nasabahRepository)
	transactionService := service.NewTransactionService(nasabahTransactionRepository, nasabahRepository)

	// NOTE: Call Handler
	nasabahHandler := handler.NewNasabahHandler(registerService, transactionService)

	// NOTE: Call Router
	routeConfig := router.NewRouter(nasabahHandler)
	routeConfig.RegisterApiRouter(e)

	// Handle graceful shutdown
	go func() {
		// Listen for interrupt or termination signals
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		sigReceived := <-sigs

		log.Printf("Received signal: %s. Shutting down gracefully...", sigReceived)

		// Attempt to stop the Echo server gracefully
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := e.Shutdown(ctx); err != nil {
			log.Fatalf("Error during server shutdown: %v", err)
		}
		log.Println("Server gracefully stopped")
	}()

	// Start the Echo server
	log.Println("Server is starting...")
	if err := e.Start(fmt.Sprintf(":%s", config.ConfigEnv.ApiPort)); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
