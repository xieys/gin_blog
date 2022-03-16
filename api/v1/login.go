package v1

import (
	"gin_blog/middleware"
	"gin_blog/model"
	"gin_blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)

	var token string
	var code int

	code = model.CheckLogin(data.Username, data.Password)

	if code == errmsg.SUCCESS {
		token, code = middleware.CreateToken(data.Username)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
		"token":  token,
	})
}
