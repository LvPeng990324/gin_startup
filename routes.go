package main

import (
	"github.com/gin-gonic/gin"

	"contract_system_server/common"
	"contract_system_server/controller"
	"contract_system_server/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	// r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	// r.POST("/api/auth/register", controller.Register)
	// r.POST("/api/auth/login", controller.Login)
	// r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	// categoryRoutes := r.Group("/categories")
	// categoryController := controller.NewCategoryController()
	// categoryRoutes.POST("", categoryController.Create)
	// categoryRoutes.PUT("/:id", categoryController.Update) //替换
	// categoryRoutes.GET("/:id", categoryController.Show)
	// categoryRoutes.DELETE("/:id", categoryController.Delete)

	// postRoutes := r.Group("/posts")
	// postRoutes.Use(middleware.AuthMiddleware())
	// postController := controller.NewPostController()
	// postRoutes.POST("", postController.Create)
	// postRoutes.PUT("/:id", postController.Update) //替换
	// postRoutes.GET("/:id", postController.Show)
	// postRoutes.DELETE("/:id", postController.Delete)
	// postRoutes.POST("/page/list", postController.PageList)

	r.Use(middleware.LogMiddleware())

	r.GET("/hello", controller.HelloWorld)
	r.GET("/hello-middleware", middleware.TestMiddleware(), controller.HelloWorld)

	userRoutes := r.Group("/user")
	userRoutes.POST("/create", controller.CreateUser)
	userRoutes.POST("/get", controller.GetUserById)
	userRoutes.POST("/get-all", controller.GetAllUser)

	redisRoutes := r.Group("/redis")
	redisRoutes.POST("/set", controller.SetKey)
	redisRoutes.POST("/get", controller.GetKey)
	redisRoutes.POST("/del", controller.DelKey)

	common.AddInfo("router init successfully!")

	return r
}
