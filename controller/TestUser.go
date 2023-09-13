package controller

import (
	"contract_system_server/dao"
	"contract_system_server/dto"
	"contract_system_server/model"
	"contract_system_server/util"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	name := ctx.PostForm("name")
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")

	user := model.User{
		Name:     name,
		Phone:    phone,
		Password: password,
	}
	err := dao.CreateUser(user)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "create user failed!",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "success!",
	})
}

func GetUserById(ctx *gin.Context) {
	idStr := ctx.PostForm("id")
	// id转为uint=
	idUint, _ := util.String2Uint(idStr)
	user, err := dao.GetUserById(idUint)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "get user failed!",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "success!",
		"data":    dto.ToUserDto(user),
	})
}

func GetAllUser(ctx *gin.Context) {
	users, err := dao.GetAllUser()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "get all user failed!",
			"error":   err.Error(),
		})
		return
	}

	// 将model.User转为dto.UserDto
	to_user_dtos := make([]dto.UserDto, 0)
	for _, user := range users {
		to_user_dtos = append(to_user_dtos, dto.ToUserDto(user))
	}

	ctx.JSON(200, gin.H{
		"message": "success!",
		"data":    to_user_dtos,
	})
}
