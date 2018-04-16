package asynlogger

import (
	"fmt"
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
	logMode        string
	logHandNoError *os.File
	logHandError   *os.File
	logChanSize    int
	LogDataChan    chan *LogData
}

func NewLogFile(config map[string]string) (log LogInterface, err error) {
	logPath, ok := config["log_path"]
	if !ok {
		err = fmt.Errorf("not find log_path ")
		return
	}

	logName, ok := config["log_name"]
	if !ok {
		err = fmt.Errorf("not find log_name ")
	}

	logMode, ok := config["log_mode"]
	if !ok {
		logMode = ""
	}

	logLevelTxt, ok := config["log_level"]
	if !ok {
		logLevelTxt = "DEBUG"
	}
	logLevelEnum := getLevelEnum(logLevelTxt)

	logChanSizeTxt, ok := config["log_chan_size"]
	if !ok {
		logChanSizeTxt = "50000"
	}
	nLogChanSize, err := strconv.Atoi(logChanSizeTxt)
	if err != nil {
		nLogChanSize = 50000
	}

	log = &LogFile{
		logLevel:    logLevelEnum,
		logPath:     logPath,
		logMode:     logMode,
		logName:     logName,
		LogDataChan: make(chan *LogData, nLogChanSize),
	}

	log.Init()

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
	go func() {
		for pLogData := range lf.LogDataChan {
			pFileHand := lf.logHandNoError
			if pLogData.BIsError {
				pFileHand = lf.logHandError
			}

			//[日期时间][日志级别][文件名；调用函数；产生日志行号][软件模块][信息内容]
			fmt.Fprintf(pFileHand, "[%s][%s][%s; %s; %d][%s][%s]\n",
				pLogData.StrTime,
				pLogData.StrLevel,
				pLogData.StrFileName, pLogData.StrFuncName, pLogData.NLineNo,
				pLogData.StrMode,
				pLogData.StrMessage,
			)
		}
	}()
}

func (lf *LogFile) SetLevel(levelEnum int) {
	if levelEnum < LogLevelDebug || levelEnum > LogLevelFatal {
		levelEnum = LogLevelDebug
	}
	lf.logLevel = levelEnum
}

/*
非核心代码异步化:
1.当业务调用打印日志的方法时，把日志相关数据写入到chan（队列），若chan满了直接丢弃日志数据
2.有一个协程不断从chan获取日志数据，最终写入文件
*/
func (lf *LogFile) writeLogToChan(level int, mode string, format string, args ...interface{}) {
	//生成日志数据类型，放入chan，若chan满了直接丢弃日志数据
	logData := createLogData(level, mode, format, args...)
	select {
	case lf.LogDataChan <- logData:
	default:
	}
}

func (lf *LogFile) Debug(format string, args ...interface{}) {
	if lf.logLevel > LogLevelDebug {
		return
	}
	lf.writeLogToChan(LogLevelDebug, lf.logMode, format, args...)
}

func (lf *LogFile) Trance(format string, args ...interface{}) {
	if lf.logLevel > LogLevelTrace {
		return
	}
	lf.writeLogToChan(LogLevelTrace, lf.logMode, format, args...)
}

func (lf *LogFile) Info(format string, args ...interface{}) {
	if lf.logLevel > LogLevelInfo {
		return
	}
	lf.writeLogToChan(LogLevelInfo, lf.logMode, format, args...)
}

func (lf *LogFile) Warn(format string, args ...interface{}) {
	if lf.logLevel > LogLevelWarn {
		return
	}
	lf.writeLogToChan(LogLevelWarn, lf.logMode, format, args...)
}

func (lf *LogFile) Error(format string, args ...interface{}) {
	if lf.logLevel > LogLevelError {
		return
	}
	lf.writeLogToChan(LogLevelError, lf.logMode, format, args...)
}

func (lf *LogFile) Fatal(format string, args ...interface{}) {
	if lf.logLevel > LogLevelFatal {
		return
	}
	lf.writeLogToChan(LogLevelFatal, lf.logMode, format, args...)
}

func (lf *LogFile) Close() {
	lf.logHandNoError.Close()
	lf.logHandError.Close()
}
