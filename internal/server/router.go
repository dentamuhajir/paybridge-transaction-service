package server

import (
	"paybridge-transaction-service/internal/health"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type Dependencies struct {
	DB *pgxpool.Pool
}

func NewRouter(deps *Dependencies) *echo.Echo {
	e := echo.New()
	healthService := health.NewService(deps.DB)
	healthHandler := health.NewHandler(*healthService)
	healthHandler.RegisterRoutes(e.Group("/api/v1"))
	return e
}
