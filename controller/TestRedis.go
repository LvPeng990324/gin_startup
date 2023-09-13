package controller

import (
	"contract_system_server/common"

	"github.com/gin-gonic/gin"
)

func SetKey(ctx *gin.Context) {
	key := ctx.PostForm("key")
	value := ctx.PostForm("value")

	err := common.Set(key, value)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "create key failed!",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "success!",
	})
}

func GetKey(ctx *gin.Context) {
	key := ctx.PostForm("key")

	value, err := common.Get(key)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "get key failed!",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "success!",
		"data":    value,
	})
}

func DelKey(ctx *gin.Context) {
	key := ctx.PostForm("key")

	err := common.Del(key)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "delete key failed!",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "success!",
	})
}
