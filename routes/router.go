package routes

import (
	v1 "gin_blog/api/v1"
	"gin_blog/middleware"
	"gin_blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)

	r := gin.Default()

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		// 文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)

		// 分类模块的用户接口
		auth.POST("category/add", v1.AddCate)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)
	}

	router := r.Group("api/v1")
	{
		// 用户模块的路由接口
		router.POST("user/add", v1.AddUser) // 用户注册
		router.GET("users", v1.GetUsers)
		router.POST("login", v1.Login) // 用户登入

		// 文章模块的路由接口
		router.GET("articles", v1.GetArticles)             // 查询文章列表
		router.GET("article/list/:id", v1.GetCateArticles) // 查询分类下的所有文章
		router.GET("article/info/:id", v1.GetArticleInfo)  // 查询单个文章信息

		// 分类模块的用户接口
		router.GET("categories", v1.GetCates)
	}

	r.Run(utils.HttpPort)
}
