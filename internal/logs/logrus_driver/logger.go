package logrusdriver

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	Log *logrus.Logger
)

func init() {
	level, err := logrus.ParseLevel("info")
	if err != nil {
		level = logrus.DebugLevel
	}

	Log = &logrus.Logger{
		Level:        level,
		Out:          os.Stdout,
		Formatter:    &logrus.JSONFormatter{},
		ReportCaller: true,
	}
}

func Info(msg string) {
	Log.Info(msg)
}

func Error(msg string, fields logrus.Fields) {
	Log.WithFields(fields).Error(msg)
}
