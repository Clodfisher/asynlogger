package asynlogger

/*
该文件主要用于日志类型的枚举定义。
通过日志种类的枚举类型获取相应字符串。
通过日志种类字符串获取到相应的枚举。
*/

const (
	LogLevelDebug = iota
	LogLevelTrace
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)

//用于如何对日志文件进行切分，0 - 通过小时进行切分， 1 - 通个大小进行切分
const (
	LogSplitTypeHour = iota
	LogSplitTypeSize
)

func getLevelText(levelEnum int) string {
	var levelText string
	switch levelEnum {
	case LogLevelDebug:
		levelText = "DEBUG"
	case LogLevelTrace:
		levelText = "TRACE"
	case LogLevelInfo:
		levelText = "INFO"
	case LogLevelWarn:
		levelText = "WARN"
	case LogLevelError:
		levelText = "ERROR"
	case LogLevelFatal:
		levelText = "FATAL"
	default:
		levelText = "DEBUG"
	}

	return levelText
}

func getLevelEnum(leveText string) int {
	var levelEnum int
	switch leveText {
	case "DEBUG":
		levelEnum = LogLevelDebug
	case "TRACE":
		levelEnum = LogLevelTrace
	case "INFO":
		levelEnum = LogLevelInfo
	case "WARN":
		levelEnum = LogLevelWarn
	case "ERROR":
		levelEnum = LogLevelError
	case "FATAL":
		levelEnum = LogLevelFatal
	default:
		levelEnum = LogLevelDebug
	}

	return levelEnum
}
