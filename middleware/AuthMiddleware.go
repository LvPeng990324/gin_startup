package middleware

import (
	"contract_system_server/common"
	"contract_system_server/dao"
	"contract_system_server/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		// 验证token
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(401, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:] // 截取字符

		token, claims, err := common.ParseToken(tokenString)

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		// token通过验证，获取claims中的UserId
		userId := claims.UserId
		var user model.User
		user, _ = dao.GetUserById(userId)

		// 验证用户是否存在
		if user.Id == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		// 用户存在，将user信息写入上下文
		ctx.Set("user", user)

		ctx.Next()
	}
}
