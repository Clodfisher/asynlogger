package asynlogger

import (
	"fmt"
	"testing"
	"time"
)

/*
	本文档用于测试logger日志库的使用情况
*/

//测试文件日志的输出
func TestFileLogger(t *testing.T) {
	config := make(map[string]string, 8)
	config["log_level"] = "DEBUG"
	config["log_mode"] = "TestLogger"
	config["log_path"] = "c:/logs/"
	config["log_name"] = "user_server"
	config["log_chan_size"] = "50000"
	err := LoggerInit("file", config)
	if err != nil {
		fmt.Println("init LoggerInit err : ", err)
	}

	Debug("init LoggerInit success...")
	Trace("Trace test %s", "you are right")
	Info("Info test %d", 10)
	Warn("Trace test %s", "you are right")
	Error("Error Test")
	Fatal("Fatal test")

	time.Sleep(time.Second * 5)

	LoggerClose()
}

//测试终端日志的输出
func TestConsoleLogger(t *testing.T) {
	config := make(map[string]string, 8)
	config["log_level"] = "DEBUG"
	config["log_mode"] = "TestLogger"
	err := LoggerInit("console", config)
	if err != nil {
		fmt.Println("init LoggerInit err : ", err)
	}

	Debug("init LoggerInit success...")
	Trace("Trace test %s", "you are right")
	Info("Info test %d", 10)
	Warn("Trace test %s", "you are right")
	Error("Error Test")
	Fatal("Fatal test")

}
