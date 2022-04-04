package logruslogger

import (
	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	entry *logrus.Entry
}

func (l *LogrusLogger) Info(message string) {
	l.entry.Info(message)
}

func (l *LogrusLogger) Error(message string) {
	l.entry.Error(message)
}

func (l *LogrusLogger) ErrorWithDetail(err error, message string) {
	l.entry.WithError(err).Error(message)
}

func NewLogrusLogger(e *logrus.Entry) *LogrusLogger {
	return &LogrusLogger{
		entry: e,
	}
}
