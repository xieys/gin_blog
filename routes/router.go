package routes

import (
	v1 "gin_blog/api/v1"
	"gin_blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)

	r := gin.Default()

	router := r.Group("api/v1")
	{
		// 用户模块的路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)

		// 文章模块的路由接口

		// 分类模块的用户接口
	}

	r.Run(utils.HttpPort)
}
