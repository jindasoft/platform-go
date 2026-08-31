package xlogger

import (
	"context"
	"os"

	"github.com/jindasoft/platform-go/xconst"
	prettyconsole "github.com/thessem/zap-prettyconsole"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Level zapcore.Level
	Name  string // application name
}

// MustInit initializes and replaces zap's global logger.
func Init(level zapcore.Level, name string, pretty bool) {
	cfg := Config{
		Level: level,
		Name:  name,
	}

	if !pretty {
		logger := defaultConfig(cfg)
		zap.ReplaceGlobals(logger)
	} else {
		log, err := createPrettyLogger(cfg)
		if err != nil {
			SysWarnf("cannot init pretty xlogger: %v; fallback to default", err)
		} else {
			zap.ReplaceGlobals(log)
		}
	}
}

func createPrettyLogger(inputCfg Config) (*zap.Logger, error) {
	// Retrieve log level from configuration
	logLevel := inputCfg.Level

	// Initialize Zap production config
	cfg := zap.NewProductionConfig()
	cfg.Development = true
	cfg.Encoding = "pretty_console" // Use pretty console encoder

	// Set the log level in the config
	cfg.Level = zap.NewAtomicLevelAt(logLevel)

	// Customize the encoder configuration for better readability
	encoderConfig := prettyconsole.NewEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeDuration = zapcore.StringDurationEncoder

	cfg.EncoderConfig = encoderConfig

	// Build and return the logger
	return cfg.Build()
}

func defaultConfig(inputCfg Config) *zap.Logger {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= inputCfg.Level
	})

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(os.Stdout), highPriority),
	)

	logger := zap.New(core, zap.AddCaller())
	logger.Named(inputCfg.Name)

	return logger
}

func getServiceNameOrDefault(ctx context.Context) string {
	str, ok := ctx.Value(xconst.ContextServiceName).(string)
	if !ok {
		return ""
	}

	return str
}

func getEnvironmentOrDefault(ctx context.Context) string {
	str, ok := ctx.Value(xconst.ContextEnvironment).(string)
	if !ok {
		return ""
	}

	return str
}
