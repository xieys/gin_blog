package v1

import (
	"gin_blog/model"
	"gin_blog/utils/errmsg"
	"gin_blog/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	var msg string
	_ = c.ShouldBindJSON(&data)

	//数据验证
	msg, code = validator.Validate(&data)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    msg,
		})
		return
	}

	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// 查询单个用户

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
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

	data, total := model.GetUsers(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"total":  total,
		"msg":    errmsg.GetErrMsg(code),
	})

}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.CheckUser(user.Username)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &user)
	}

	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})

}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}
