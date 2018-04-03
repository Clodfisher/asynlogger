package asynlogger

import "os"

/*
该文件主要用于将日志打印到文件中
*/

type LogFile struct {
	logLevel        int
	filePath        string
	fileName        string
	fileHandNoError os.File
	fileHandError   os.File
	LogDataChan     chan *LogData
}

func NewLogFile(config map[string]string) (LogInterface, error) {

}

func (lf *LogFile) Init() {

}

func (lf *LogFile) SetLevel(levelEnum int) {

}

func (lf *LogFile) Debug(format string, args ...interface{}) {

}

func (lf *LogFile) Trance(format string, args ...interface{}) {

}

func (lf *LogFile) Info(format string, args ...interface{}) {

}

func (lf *LogFile) Warn(format string, args ...interface{}) {

}

func (lf *LogFile) Error(format string, args ...interface{}) {

}

func (lf *LogFile) Fatal(format string, args ...interface{}) {

}

func (lf *LogFile) Close() {

}
