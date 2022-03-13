package v1

import (
	"gin_blog/model"
	"gin_blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddCate 添加分类
func AddCate(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)

	code = model.CheckCate(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateCate(&data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// GetCates 查询分类列表
func GetCates(c *gin.Context) {
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

	data := model.GetCates(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})

}

// todo 查询分类下的所有文章

// EditCate 编辑分类
func EditCate(c *gin.Context) {
	var cate model.Category
	_ = c.ShouldBindJSON(&cate)
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.CheckCate(cate.Name)
	if code == errmsg.SUCCESS {
		model.EditCate(id, &cate)
	}

	if code == errmsg.ERROR_CATEGORY_USED {
		c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})

}

// DeleteCate 删除分类
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCate(id)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}
