package log

import (
	"go.uber.org/zap"
	l "log"
)

var production *zap.Logger

func InitLogger() {
	logger, err := zap.NewProduction()
	if err != nil {
		l.Fatalf("can't initialize zap logger: %v", err)
	}

	production = logger
}

func Logger() *zap.Logger {
	return production
}
