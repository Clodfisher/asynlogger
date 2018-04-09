package asynlogger

import (
	"fmt"
	"strconv"
)

/*
该文件主要用于将日志打印控制台
对于控制台日志打印，是要求实时性，只要在后台调试时查看的信息，因此不用开启并发处理
*/

type LogConsole struct {
	LogLevel int
}

func NewLogConsole(config map[string]string) (log LogInterface, err error) {

	logLevelTxt, ok := config["log_level"]
	if ok {
		logLevelTxt = "DEBUG"
	}
	logLevelEnum := getLevelEnum(logLevelTxt)

	log = &LogConsole{
		LogLevel: logLevelEnum,
	}

	return
}

func (lc *LogConsole) Init() {

}

func (lc *LogConsole) SetLevel(levelEnum int) {
	if levelEnum < LogLevelDebug || levelEnum > LogLevelFatal {
		levelEnum = LogLevelDebug
	}

	lc.LogLevel = levelEnum
}

func (lc *LogConsole) writeLogToConsole(level int, format string, args ...interface{}) {
	//生成日志数据类型，放入chan，若chan满了直接丢弃日志数据
	pLogData := createLogData(level, format, args)

	//[日期时间][日志级别][文件名；调用函数；产生日志行号][用户ID][软件模块][信息内容]
	fmt.Printf("[%s][%s][%s; %s; %d][%s]\n",
		pLogData.StrTime,
		pLogData.StrLevel,
		pLogData.StrFileName, pLogData.StrFuncName, pLogData.NLineNo,
		pLogData.StrMessage,
	)

}

func (lc *LogConsole) Debug(format string, args ...interface{}) {
	if lc.LogLevel > LogLevelDebug {
		return
	}

	lc.writeLogToConsole(LogLevelDebug, format, args)
}

func (lc *LogConsole) Trance(format string, args ...interface{}) {
	if lc.LogLevel > LogLevelTrace {
		return
	}

	lc.writeLogToConsole(LogLevelTrace, format, args)
}

func (lc *LogConsole) Info(format string, args ...interface{}) {
	if lc.LogLevel > LogLevelInfo {
		return
	}

	lc.writeLogToConsole(LogLevelInfo, format, args)
}

func (lc *LogConsole) Warn(format string, args ...interface{}) {
	if lc.LogLevel > LogLevelWarn {
		return
	}

	lc.writeLogToConsole(LogLevelWarn, format, args)
}

func (lc *LogConsole) Error(format string, args ...interface{}) {
	if lc.LogLevel > LogLevelError {
		return
	}

	lc.writeLogToConsole(LogLevelError, format, args)
}

func (lc *LogConsole) Fatal(format string, args ...interface{}) {
	if lc.LogLevel > LogLevelFatal {
		return
	}

	lc.writeLogToConsole(LogLevelFatal, format, args)
}

func (lc *LogConsole) Close() {

}
