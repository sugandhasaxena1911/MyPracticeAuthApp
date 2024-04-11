package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func init() {
	var err error
	//Log, err = zap.NewProduction(zap.AddCallerSkip(1))
	// add custom configuration
	encoderconfig := zap.NewProductionEncoderConfig()
	encoderconfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderconfig.TimeKey = "timestamp"
	config := zap.NewProductionConfig()
	config.EncoderConfig = encoderconfig
	Log, err = config.Build()
	if err != nil {
		panic(err)
	}

}

// you should abstract the Log variable & create a function , so that even in future you start using some
// other logger , you can easily do it without changing logging that you did at every place in application
// but the trace will show "caller":"logger/logger.go:19"  for all .
// Zap has a solution for that also
// change the json tags & time unit
func Info(message string, fields ...zap.Field) {
	Log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	Log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	Log.Error(message, fields...)
}
