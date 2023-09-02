package logging

import (
	"github.com/gost1k337/url_shortener/api_gateway_service/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger(cfg *config.Config) Logger {
	zapConfig := zap.NewDevelopmentConfig()

	if cfg.App.Debug {
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	if !cfg.App.Debug && cfg.App.LogOutput != "" {
		zapConfig.OutputPaths = []string{cfg.App.LogOutput}
	}

	logger, _ := zapConfig.Build()
	sugar := logger.Sugar()
	return Logger{sugar}
}
