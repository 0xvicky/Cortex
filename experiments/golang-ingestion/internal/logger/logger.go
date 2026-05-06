package logger

import "go.uber.org/zap"

var Log *zap.Logger

func InitLogger() {
	var logErr error
	Log, logErr = zap.NewDevelopment()
	if logErr != nil {
		panic(logErr) // fail fast, don’t silently continue
	}
}
