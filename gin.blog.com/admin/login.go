package controller

import (
	"gin.blog.com/common"
	"gin.blog.com/handler"
	"gin.blog.com/models"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type auth struct {
	UserName string	`valid:"Required; MaxSize(50)"`
	PassWord string	`valid:"Required; MaxSize(50)"`
}



func Login(c *gin.Context)  {
	//结构体绑定json 后解析
	from := models.LoginInfo{}
	//解析json
	c.BindJSON(&from)
	username := from.Username
	password := from.Password
	valid := validation.Validation{}
	a := auth{UserName: username, PassWord: password}
	ok,_ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := common.INVALID_PARAMS

	if ok {
		isExits := models.CheckAuth(username, password)

		if isExits { //验证数据是正确
			token, err := common.GenerateToken(username, password) // 生成token
			if err != nil { //生成失败
				code = common.ERROR_AUTH_TOKEN
			}else{
				nowTime := time.Now()
				expireTime := nowTime.Add(3 * time.Hour)
				data["token"] = token
				data["user"] = username
				data["expireAt"] = expireTime

				m := make(map[string]interface{})
				m["id"] = "admin"
				a := make(map[int]interface{})
				a[0] = m
				roles := make(map[string]interface{})
				roles["id"] = "admin"
				roles["operation"] = a
				data["roles"] = a
				code = common.SUCCESS

			}
		}else{ //密码错误
			code = common.ERROR_AUTH
		}
	}else{
		//展示验证规则遍历
		for _, err := range valid.Errors{
			log.Println(err.Key, err.Message)
		}
	}
	//返回json处理
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"mgs"  : common.GetMsg(code),
		"data" : data,
	})
}

func Info(c *gin.Context) {
	username := handler.Claims.Username
	data := models.GetUserOne(username)

	c.JSON(http.StatusOK, gin.H{
		"code" : 200,
		"mgs"  : "成功",
		"data" : data,
	})
}

func Routers(c *gin.Context) {
	data := make(map[string]interface{})
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"mgs":  "成功",
		"data": data,
	})
}