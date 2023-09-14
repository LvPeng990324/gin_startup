package common

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func InitLog() {
	// 关闭gin默认的日志
	gin.DefaultWriter = io.Discard

	// 设置日志格式
	debug_flag := viper.GetBool("server.debug")
	if debug_flag {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	// 设置time字段格式化风格
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix  // 算了，默认就挺好，需要的话可以设置成别的，比如时间戳

	// 配置输出到文件
	// 创建记录日志的软链接文件
	path := viper.GetString("server.log_files")
	log_clear_days := viper.GetInt("server.log_clear_days")
	writer, _ := rotatelogs.New(
		path+"%Y%m%d.log", // 每天分日志文件
		rotatelogs.WithLinkName(path),
		// 日志文件保存时长
		rotatelogs.WithMaxAge(time.Duration(24*log_clear_days)*time.Hour),
		//这里设置1分钟检查一次是否要分日志文件
		rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
	)
	log.Logger = log.Output(writer)
}

// 获取logger
func get_logger(log_level string) *zerolog.Event {
	if log_level == "debug" {
		return log.Debug()
	} else if log_level == "info" {
		return log.Info()
	} else if log_level == "warn" {
		return log.Warn()
	} else {
		return log.Error()
	}
}

// 基础的加日志以及对应的快捷日志方法，用于手动调用
func add_log(log_level string, msg string) {
	logger := get_logger(log_level)
	logger.Msg(msg)
}
func AddDebug(msg string) {
	add_log("debug", msg)
}
func AddInfo(msg string) {
	add_log("info", msg)
}
func AddWarn(msg string) {
	add_log("warn", msg)
}
func AddError(msg string) {
	add_log("error", msg)
}

// middleware用日志方法
func MiddlewareLog(ip string, host string, url string, status int, delay int64) {
	logger := get_logger("info")
	logger.
		Str("ip", ip).
		Str("host", host).
		Str("url", url).
		Int("status", status).
		Int64("delay", delay).
		Send()
}
