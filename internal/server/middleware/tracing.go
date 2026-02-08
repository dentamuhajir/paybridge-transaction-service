package middleware

// import (
// 	"context"
// 	"paybridge-transaction-service/internal/infra/logger"

// 	"github.com/labstack/echo/v4"
// )

// // Middleware to generate or propagate trace_id and span_id
// func TracingMiddleware() echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			req := c.Request()

// 			// Extract trace_id from headers if present
// 			traceID := req.Header.Get("X-Trace-Id")
// 			if traceID == "" {
// 				traceID = logger.NewUUID()
// 			}

// 			// Generate new span_id for this service
// 			spanID := logger.NewUUID()

// 			// Store in context
// 			ctx := context.WithValue(req.Context(), logger.CtxTraceID, traceID)
// 			ctx = context.WithValue(ctx, logger.CtxSpanID, spanID)
// 			c.SetRequest(req.WithContext(ctx))

// 			// Call next handler
// 			return next(c)
// 		}
// 	}
// }
