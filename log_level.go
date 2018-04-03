package asynlogger

/*
该文件主要用于日志类型的枚举定义。
通过日志种类的枚举类型获取相应字符串。
通过日志种类字符串获取到相应的枚举。
*/

const (
	LogLevelDebug = iota
	LogLevelTrance
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)

func getLevelText(levelEnum int) string {
	var levelText string
	switch levelEnum {
	case LogLevelDebug:
		levelText = "DEBUG"
	case LogLevelTrance:
		levelText = "TRANCE"
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
	case "TRANCE":
		levelEnum = LogLevelTrance
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
