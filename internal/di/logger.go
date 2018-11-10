package di

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func ProvideLogger() *logrus.Logger {
	if logger != nil {
		return logger
	}

	logger = logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)

	return logger
}
