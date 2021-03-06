package v1

import (
	"gin_blog/model"
	"gin_blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddArticle 添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)

	code = model.CreateArt(&data)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// GetCateArticles 查询分类下的所有文章
func GetCateArticles(c *gin.Context) {
	var artList []model.Article
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 10:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	var total int64
	artList, code, total = model.GetCateArts(id, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   artList,
		"total":  total,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// GetArticleInfo 查询单个文章
func GetArticleInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArtInfo(id)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})

}

// GetArticles 查询文章列表
func GetArticles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 10:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}

	data, code, total := model.GetArts(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"total":  total,
		"msg":    errmsg.GetErrMsg(code),
	})

}

// EditArticle 编辑文章
func EditArticle(c *gin.Context) {
	var art model.Article
	_ = c.ShouldBindJSON(&art)
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.EditArt(id, &art)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})

}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArt(id)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}
