package asynlogger

import (
	"fmt"
)

/*
该文件主要用于将日志打印控制台
对于控制台日志打印，是要求实时性，只要在后台调试时查看的信息，因此不用开启并发处理
*/

type LogConsole struct {
	logLevel int
	logMode  string
}

func NewLogConsole(config map[string]string) (log LogInterface, err error) {

	logLevelTxt, ok := config["log_level"]
	if !ok {
		logLevelTxt = "DEBUG"
	}
	logLevelEnum := getLevelEnum(logLevelTxt)

	logMode, ok := config["log_mode"]
	if !ok {
		logMode = ""
	}

	log = &LogConsole{
		logLevel: logLevelEnum,
		logMode:  logMode,
	}

	return
}

func (lc *LogConsole) Init() {

}

func (lc *LogConsole) SetLevel(levelEnum int) {
	if levelEnum < LogLevelDebug || levelEnum > LogLevelFatal {
		levelEnum = LogLevelDebug
	}

	lc.logLevel = levelEnum
}

func (lc *LogConsole) writeLogToConsole(level int, mode string, format string, args ...interface{}) {
	//生成日志数据类型，放入chan，若chan满了直接丢弃日志数据
	pLogData := createLogData(level, mode, format, args...)

	//[日期时间][日志级别][文件名；调用函数；产生日志行号][软件模块][信息内容]
	fmt.Printf("[%s][%s][%s; %s; %d][%s][%s]\n",
		pLogData.StrTime,
		pLogData.StrLevel,
		pLogData.StrFileName, pLogData.StrFuncName, pLogData.NLineNo,
		pLogData.StrMode,
		pLogData.StrMessage,
	)

}

func (lc *LogConsole) Debug(format string, args ...interface{}) {
	if lc.logLevel > LogLevelDebug {
		return
	}

	lc.writeLogToConsole(LogLevelDebug, lc.logMode, format, args...)
}

func (lc *LogConsole) Trance(format string, args ...interface{}) {
	if lc.logLevel > LogLevelTrace {
		return
	}

	lc.writeLogToConsole(LogLevelTrace, lc.logMode, format, args...)
}

func (lc *LogConsole) Info(format string, args ...interface{}) {
	if lc.logLevel > LogLevelInfo {
		return
	}

	lc.writeLogToConsole(LogLevelInfo, lc.logMode, format, args...)
}

func (lc *LogConsole) Warn(format string, args ...interface{}) {
	if lc.logLevel > LogLevelWarn {
		return
	}

	lc.writeLogToConsole(LogLevelWarn, lc.logMode, format, args...)
}

func (lc *LogConsole) Error(format string, args ...interface{}) {
	if lc.logLevel > LogLevelError {
		return
	}

	lc.writeLogToConsole(LogLevelError, lc.logMode, format, args...)
}

func (lc *LogConsole) Fatal(format string, args ...interface{}) {
	if lc.logLevel > LogLevelFatal {
		return
	}

	lc.writeLogToConsole(LogLevelFatal, lc.logMode, format, args...)
}

func (lc *LogConsole) Close() {

}
