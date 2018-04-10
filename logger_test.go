package asynlogger

import (
	"testing"
)

/*
	本文档用于测试logger日志库的使用情况
*/

//测试文件日志的输出
func TestFileLogger(t *testing.T) {

}

//测试终端日志的输出
func TestConsoleLogger(t *testing.T) {
	config := make(map[string]string, 8)
	config["log_level"] = "DEBUG"
	config["log_mode"] = "TestLogger"
	err := LoggerInit("console", config)
	if err != nil {
		println("init LoggerInit err : ", err)
	}

	Debug("init LoggerInit success...")
	Trace("Trace test %s\n", "you are right")
	Info("Info test %d\n", 10)
	Warn("Trace test %s\n", "you are right")
	Error("Error Test")
	Fatal("Fatal test\n")

}
