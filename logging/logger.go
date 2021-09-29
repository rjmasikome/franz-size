package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger creates a preconfigured global logger and configures the global zap logger
func NewLogger(cfg Config) *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	// Parse log level text to zap.LogLevel. Error check isn't required because the input is already validated.
	level := zap.NewAtomicLevel()
	_ = level.UnmarshalText([]byte(cfg.Level))

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		level,
	)
	logger := zap.New(core)
	zap.ReplaceGlobals(logger)

	return logger
}
