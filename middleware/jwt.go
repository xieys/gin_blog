package middleware

import (
	"gin_blog/utils"
	"gin_blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var jwtKey = []byte(utils.JwtKey)

// CreateToken 生成token
func CreateToken(username string) (string, int) {
	claims := MyClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Hour)),
			Issuer:    "blog",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return ss, errmsg.SUCCESS
}

// ParserToken 解析token
func ParserToken(tokenString string) (*MyClaims, int) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Is(jwt.ErrTokenExpired) {
				return nil, errmsg.ERROR_TOKEN_RUNTIME
			} else {
				return nil, errmsg.ERROR_TOKEN_WRONG
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, errmsg.SUCCESS
		}
		return nil, errmsg.ERROR_TOKEN_WRONG
	}
	return nil, errmsg.ERROR_TOKEN_WRONG
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		// 解析token
		claims, code := ParserToken(checkToken[1])
		if code == errmsg.ERROR_TOKEN_RUNTIME {
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
				"data":    nil,
			})
			c.Abort()
			return
		}
		if code == errmsg.ERROR_TOKEN_WRONG {
			// 其他错误
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
				"data":    nil,
			})
			c.Abort()
			return
		}

		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", claims)
		c.Next()
	}
}
