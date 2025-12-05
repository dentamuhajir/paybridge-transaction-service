package wallet

import (
	"net/http"
	"paybridge-transaction-service/internal/server/middleware"
	"paybridge-transaction-service/pkg/response"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{service: svc}
}

func (h *Handler) RegisterRoutes(g *echo.Group) {
	g.POST("/wallet", h.Create, middleware.ValidateInternalToken)
}

// CreateWalletHandler godoc
// @Summary Create wallet
// @Description Create a new wallet for a user
// @Tags Wallet
// @Accept json
// @Produce json
// @Param request body CreateWalletReq true "Wallet creation payload"
// @Success 200 {object} response.SwaggerSuccessResponse
// @Failure 400 {object} response.SwaggerErrorResponse
// @Failure 500 {object} response.SwaggerErrorResponse
// @Security InternalTokenAuth
// @Router /wallet [post]
func (h *Handler) Create(c echo.Context) error {
	var req CreateWalletReq

	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			response.Error("invalid request body", http.StatusBadRequest),
		)
	}

	resp, err := h.service.CreateWallet(c.Request().Context(), req)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			response.Error("failed to create wallet", http.StatusInternalServerError),
		)
	}

	return c.JSON(http.StatusOK,
		response.Success("wallet created", resp, http.StatusOK),
	)
}
