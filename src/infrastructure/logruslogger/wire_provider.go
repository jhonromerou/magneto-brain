package logruslogger

import (
	"github.com/jhonromerou/magneto-brain/src/domain"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

var SetLogrusLogger = wire.NewSet(
	NewLogrusLoggerProvider,
	logrus.NewEntry,
	NewLogrusLogger,
	wire.Bind(new(domain.Logger), new(*LogrusLogger)),
)

// NewLogrusLoggerProvider defines and configs a new instance of logger.
func NewLogrusLoggerProvider() *logrus.Logger {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "log.level",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "function.name",
		},
	})

	logger.SetLevel(logrus.DebugLevel)

	return logger
}
