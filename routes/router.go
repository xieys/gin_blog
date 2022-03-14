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
		router.POST("article/add", v1.AddArticle)
		router.GET("articles", v1.GetArticles)             // 查询文章列表
		router.GET("article/list/:id", v1.GetCateArticles) // 查询分类下的所有文章
		router.GET("article/info/:id", v1.GetArticleInfo)  // 查询单个文章信息
		router.PUT("article/:id", v1.EditArticle)
		router.DELETE("article/:id", v1.DeleteArticle)

		// 分类模块的用户接口
		router.POST("category/add", v1.AddCate)
		router.GET("categories", v1.GetCates)
		router.PUT("category/:id", v1.EditCate)
		router.DELETE("category/:id", v1.DeleteCate)
	}

	r.Run(utils.HttpPort)
}
