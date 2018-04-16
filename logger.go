package asynlogger

import (
	"fmt"
)

var log LogInterface

/*
	将日志库进行易用性封装：
	通过初始化函数，进行日志库对象的创建：file, "创建一个文件日志实例" console, "创建一个console日志实例"。
	通过设置日志产生等级函数，可以控制程序输出的日志，使调用者只查需要的日志级别。
	对于不同等级日志生成函数进行封装，使调用者根据此日志输出等级，输出相应等级的日志。
*/

/* @brief			日志实例初始化
 * @details
 * @param[in] 		nameLogger		根据不同的名字创建不同的日志实例
 *									file, "创建一个文件日志实例" console, "创建一个console日志实例"。
 * @param[in]		config			用于存储日志创建实例时的基本配置
 *								    config["log_path"]  -- 用于配置产生日志的路径，必须配置否则初始化失败
 *                                  config["log_name"]	-- 用于配置产生日志的名字，必须配置否则初始化失败
 *                                  config["log_mode"]	-- 用于配置产生日志模块名，默认为""
 *                                  config["log_level"]	-- 用于配置产生日志的等级，默认为"DEBUG"
 *                                  config["log_chan_size"]	-- 用于配置chan队列的大小，默认为"50000"
 *									config["log_split_type"] -- 用于配置日志文件备份迁移类型 "hour" - 按小时， size - 根据大小
 * 									config["log_split_size"] -- 用于配置当文件大小大于此值时,单位为B，将进行配置，默认为100M(104857600)
 */
func LoggerInit(nameLogger string, config map[string]string) (err error) {
	switch nameLogger {
	case "file":
		log, err = NewLogFile(config)
	case "console":
		log, err = NewLogConsole(config)
	default:
		err = fmt.Errorf("unsupport logger name:%s", nameLogger)
	}

	return
}

/* @brief			关闭日志
 * @details			进行日志输出的善后工作
 * @param[in]
 */
func LoggerClose() {
	log.Close()
}

/* @brief			设置显示日志输出的级别
 * @details			只有大于等于该等级的日志才能数据，比如设置的级别为LogLevelTrace，
 *                  则LogLevelDebug级别的日志将不会输出
 * @param[in] 		levelEnum		可以设置的显示日志级别：
 *					LogLevelDebug = iota  0 - Debug级别：用于调试程序，日志最为详细，对于程序的性能影响比较大。非错误性。
 *					LogLevelTrace		  1 - Trace级别：用于追踪问题。非错误性。
 *					LogLevelInfo		  2 - Info级别：打印程序运行过程中比较重要信息，比如访问日志。非错误性。
 *					LogLevelWarn		  3 - Warn级别：警告日志，说明程序运行出现潜在的问题。错误性。
 *					LogLevelError         4 - Error级别：错误日志，程序运行发生错误，但不影响程序的运行。错误性。
 * 					LogLevelFatal         5 - Fatal级别：严重错误，发生的错误会导致程序退出。错误性。
 */
func LoggerSetLevelShow(levelEnum int) {
	log.SetLevel(levelEnum)
}

/* @brief			输出的日志级别为Debug级别
 * @details			传入参数方式和fmt.Printf一样
 *
 * @param[in] 		format		格式化输出日志的样式
 * @param[in] 		args		格式化中需要加入的参数
 */
func Debug(format string, args ...interface{}) {
	log.Debug(format, args...)
}

/* @brief			输出的日志级别为Trace级别
 * @details			传入参数方式和fmt.Printf一样
 *
 * @param[in] 		format		格式化输出日志的样式
 * @param[in] 		args		格式化中需要加入的参数
 */
func Trace(format string, args ...interface{}) {
	log.Trance(format, args...)
}

/* @brief			输出的日志级别为Info级别
 * @details			传入参数方式和fmt.Printf一样
 *
 * @param[in] 		format		格式化输出日志的样式
 * @param[in] 		args		格式化中需要加入的参数
 */
func Info(format string, args ...interface{}) {
	log.Info(format, args...)
}

/* @brief			输出的日志级别为Warn级别
* @details			传入参数方式和fmt.Printf一样
*
* @param[in] 		format		格式化输出日志的样式
* @param[in] 		args		格式化中需要加入的参数
 */
func Warn(format string, args ...interface{}) {
	log.Warn(format, args...)
}

/* @brief			输出的日志级别为Error级别
* @details			传入参数方式和fmt.Printf一样
*
* @param[in] 		format		格式化输出日志的样式
* @param[in] 		args		格式化中需要加入的参数
 */
func Error(format string, args ...interface{}) {
	log.Error(format, args...)
}

/* @brief			输出的日志级别为Fatal级别
 * @details			传入参数方式和fmt.Printf一样
 *
 * @param[in] 		format		格式化输出日志的样式
 * @param[in] 		args		格式化中需要加入的参数
 */
func Fatal(format string, args ...interface{}) {
	log.Fatal(format, args...)
}
