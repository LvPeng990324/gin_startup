package middleware

import (
	"contract_system_server/common"
	"time"

	"github.com/gin-gonic/gin"
)

func LogMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start_timestamp := time.Now().UnixMilli()
		ctx.Next()
		end_timestamp := time.Now().UnixMilli()
		used_time := end_timestamp - start_timestamp // 算出消息处理用了多久

		ip := ctx.ClientIP()
		host := ctx.Request.Host
		url := ctx.Request.URL.String()
		status := ctx.Writer.Status()

		common.MiddlewareLog(ip, host, url, status, used_time)
	}
}
