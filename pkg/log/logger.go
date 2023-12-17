package log

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/wisnunu254/api-auth-golang/config"
)

type Logger struct {
	*logrus.Logger
}

var logger = &Logger{}

func GetLogger() *Logger {
	return logger
}

func Loggers() {
	logger = &Logger{logrus.New()}

	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetOutput(os.Stdout)
	if config.ConfigApp().Debug {
		logger.SetLevel(logrus.DebugLevel)
	}
}
