package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func TestMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("TestMiddleware in")

		ctx.Next()
	}
}
