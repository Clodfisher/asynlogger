package asynlogger

import (
	"fmt"
	"gopcp.v2/helper/log/base"
	"os"
	"strconv"
)

/*
该文件主要用于将日志打印到文件中
*/

type LogFile struct {
	logLevel       int
	logPath        string
	logName        string
	logHandNoError *os.File
	logHandError   *os.File
	logChanSize    int
	LogDataChan    chan *LogData
}

func NewLogFile(config map[string]string) (log LogInterface, err error) {
	logPath, ok := config["log_path"]
	if ok {
		err = fmt.Errorf("not find log_path ")
		return
	}

	logName, ok := config["log_name"]
	if ok {
		err = fmt.Errorf("not find log_name ")
	}

	logLevelTxt, ok := config["log_level"]
	if ok {
		logLevelTxt = "DEBUG"
	}
	logLevelEnum := getLevelEnum(logLevelTxt)

	logChanSizeTxt, ok := config["log_chan_size"]
	if ok {
		logChanSizeTxt = "50000"
	}
	nLogChanSize, err := strconv.Atoi(logChanSizeTxt)
	if err != nil {
		nLogChanSize = 50000
	}

	log = &LogFile{
		logLevel:    logLevelEnum,
		logPath:     logPath,
		logName:     logName,
		LogDataChan: make(chan *LogData, nLogChanSize),
	}

	return
}

func (lf *LogFile) Init() {

	//写非错误性日志：Debug级别、Trace级别、Info级别
	filePathName := fmt.Sprintf("%s/%s.log", lf.logPath, lf.logName)
	logHandNoError, err := os.OpenFile(filePathName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open faile %s failed, err:%v", filePathName, err))
	}
	lf.logHandNoError = logHandNoError

	//写非错误性日志：Warn级别、Error级别、Fatal级别
	filePathName = fmt.Sprintf("%s/%s.log.wef", lf.logPath, lf.logName)
	logHandError, err := os.OpenFile(filePathName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open faile %s failed, err:%v", filePathName, err))
	}
	lf.logHandError = logHandError

	//开个协程将日志写入到文件

}

func (lf *LogFile) SetLevel(levelEnum int) {
	if levelEnum < LogLevelDebug || levelEnum > LogLevelFatal {
		levelEnum = LogLevelDebug
	}
	lf.logLevel = levelEnum
}

func (lf *LogFile) writeLogToChan(level int, format string, args ...interface{}) {
	//生成日志数据类型，放入chan，若chan满了直接丢弃日志数据
	logData := createLogData(level, format, args)
	select {
	case lf.LogDataChan <- logData:
	default:
	}
}

func (lf *LogFile) Debug(format string, args ...interface{}) {
	if lf.logLevel > LogLevelDebug {
		return
	}
	lf.writeLogToChan(LogLevelDebug, format, args)
}

func (lf *LogFile) Trance(format string, args ...interface{}) {
	if lf.logLevel > LogLevelTrace {
		return
	}
	lf.writeLogToChan(LogLevelTrace, format, args)
}

func (lf *LogFile) Info(format string, args ...interface{}) {
	if lf.logLevel > LogLevelInfo {
		return
	}
	lf.writeLogToChan(LogLevelInfo, format, args)
}

func (lf *LogFile) Warn(format string, args ...interface{}) {
	if lf.logLevel > LogLevelWarn {
		return
	}
	lf.writeLogToChan(LogLevelWarn, format, args)
}

func (lf *LogFile) Error(format string, args ...interface{}) {
	if lf.logLevel > LogLevelError {
		return
	}
	lf.writeLogToChan(LogLevelError, format, args)
}

func (lf *LogFile) Fatal(format string, args ...interface{}) {
	if lf.logLevel > LogLevelFatal {
		return
	}
	lf.writeLogToChan(LogLevelFatal, format, args)
}

func (lf *LogFile) Close() {
	lf.logHandNoError.Close()
	lf.logHandError.Close()
}
