package logger

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

// Context keys for trace
type ctxKey string

const (
	CtxTraceID ctxKey = "trace_id"
	CtxSpanID  ctxKey = "span_id"
)

// New creates a zap logger instance (DI entry point)
func New() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()

	// Optional: customize encoder for ECS-like output
	cfg.EncoderConfig.TimeKey = "@timestamp"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	return cfg.Build()
}

// Info logs an info-level message with trace/span from context
func Info(ctx context.Context, message string, fields ...zap.Field) {
	traceID, _ := ctx.Value(CtxTraceID).(string)
	spanID, _ := ctx.Value(CtxSpanID).(string)

	coreFields := []zap.Field{
		zap.String("@timestamp", time.Now().UTC().Format(time.RFC3339Nano)),
		zap.String("@version", "1"),
		zap.String("logger_name", "transaction-service"),
		zap.String("thread_name", "goroutine"),
		zap.String("level", "INFO"),
		zap.Int("level_value", 20000),
		zap.String("trace_id", traceID),
		zap.String("span_id", spanID),
	}
	coreFields = append(coreFields, fields...)

	Log.Info(message, coreFields...)
}

// Error logs error-level messages
func Error(ctx context.Context, message string, err error, fields ...zap.Field) {
	traceID, _ := ctx.Value(CtxTraceID).(string)
	spanID, _ := ctx.Value(CtxSpanID).(string)

	coreFields := []zap.Field{
		zap.String("@timestamp", time.Now().UTC().Format(time.RFC3339Nano)),
		zap.String("@version", "1"),
		zap.String("logger_name", "transaction-service"),
		zap.String("thread_name", "goroutine"),
		zap.String("level", "ERROR"),
		zap.Int("level_value", 40000),
		zap.String("trace_id", traceID),
		zap.String("span_id", spanID),
		zap.Error(err),
	}
	coreFields = append(coreFields, fields...)

	Log.Error(message, coreFields...)
}

// Helper to generate new UUID string
func NewUUID() string {
	return uuid.New().String()
}
