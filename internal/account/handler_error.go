package account

import (
	"errors"
	"net/http"

	"paybridge-transaction-service/internal/infra/logger"
	"paybridge-transaction-service/pkg/response"

	"github.com/labstack/echo/v4"
)

func mapError(c echo.Context, log *logger.Logger, err error) error {
	switch {
	case errors.Is(err, ErrInvalidUserID):
		return c.JSON(http.StatusBadRequest,
			response.Error("invalid user id", http.StatusBadRequest))

	case errors.Is(err, ErrAccountNotFound):
		return c.JSON(http.StatusNotFound,
			response.Error("account not found", http.StatusNotFound))

	case errors.Is(err, ErrAccountInactive):
		return c.JSON(http.StatusConflict,
			response.Error("account inactive", http.StatusConflict))

	default:
		return c.JSON(http.StatusInternalServerError,
			response.Error("internal server error", http.StatusInternalServerError))
	}
}
