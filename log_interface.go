package asynlogger

type LogInterface interface {
	Init()
	SetLevel(levelEnum int)
	Debug(format string, args ...interface{})
	Trance(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}
