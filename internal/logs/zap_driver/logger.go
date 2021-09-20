package zapdriver

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Log *zap.Logger
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout", "logs/zap.log"},
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "level",
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
			CallerKey:    "caller",
		},
	}

	var err error
	Log, err = logConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Field(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

func Info(format string, a interface{}, tags ...zap.Field) {

	if a != nil {
		format = fmt.Sprintf(format, a)
	}

	Log.Info(format, tags...)
	Log.Sync()
}

func Error(format string, a interface{}, tags ...zap.Field) {

	if a != nil {
		format = fmt.Sprintf(format, a)
	}

	Log.Error(format, tags...)
	Log.Sync()
}
