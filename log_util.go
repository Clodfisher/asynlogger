package asynlogger

import (
	"fmt"
	"runtime"
	"time"
)

/*
该文件主要用于实现日志打印的通用工具
*/

type LogData struct {
	StrLevel    string
	StrMessage  string
	StrTime     string
	StrFileName string
	StrFuncName string
	NLineNo     int
	BIsError    bool
}

//取得行信息
func GetLineInfo() (strFileName string, strFuncName string, nLineNo int) {
	pc, file, line, ok := runtime.Caller(4)
	if ok {
		strFileName = file
		strFuncName = runtime.FuncForPC(pc).Name()
		nLineNo = line
	}

	return
}

//创建logData
func createLogData(level int, format string, args ...interface{}) *LogData {
	nowTime := time.Now()
	strTime := nowTime.Format("2006-01-02 15:04:05.999")

	levelTxt := getLevelText(level)

	strFileName, strFuncName, nLineNo := GetLineInfo()

	msg := fmt.Sprintf(format, args)

	bIsError := false
	if level == LogLevelError || level == LogLevelWarn || level == LogLevelFatal {
		bIsError = true
	}

	return &LogData{
		StrLevel:    levelTxt,
		StrMessage:  msg,
		StrTime:     strTime,
		StrFileName: strFileName,
		StrFuncName: strFileName,
		NLineNo:     nLineNo,
		BIsError:    bIsError,
	}
}
