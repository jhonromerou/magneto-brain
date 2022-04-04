package domain

type Logger interface {
	Error(message string)
	ErrorWithDetail(err error, message string)
	Info(message string)
}
