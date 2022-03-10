package main

import (
	"gin_blog/model"
	"gin_blog/routes"
)

func main() {

	//初始化数据库
	model.InitDb()

	// 初始化路由
	routes.InitRouter()
}
