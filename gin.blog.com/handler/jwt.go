package handler

import (
	"fmt"
	"gin.blog.com/common"
	"gin.blog.com/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var Claims models.LoginInfo

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = common.SUCCESS //回传状态码
		token := c.Query("token")
		fmt.Println(token)

		if token == "" {
			token = c.GetHeader("Authorization")
		}
		if token == "" {
			code = common.INVALID_PARAMS
		} else {
			// tooken 解析
			claims, err := common.ParseToken(token)

			Claims.Username = claims.Username
			Claims.Password = claims.Password
			if err != nil {
				//t解析错误
				code = common.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt{ //当前时间大于解析出来的token时间
				code = common.ERROR_AUTH_CHECK_TOKEN_TIMEOUT //token 过期
			}
		}

		// 验证失败 返回错误信息
		if code != common.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg"  : common.GetMsg(code),
				"data" : data,
			})

			c.Abort()
			return
		}

		//jwt 中间件验证下一步
		c.Next()
	}
}

