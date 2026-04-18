package logger

import "go.uber.org/zap"

var Log *zap.Logger

func InitLogger() *zap.Logger {
	Log, logErr := zap.NewProduction()
	if logErr != nil {
		panic(logErr) // fail fast, don’t silently continue
	}
	return Log
}
