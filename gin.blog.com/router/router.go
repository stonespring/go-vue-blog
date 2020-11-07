package router

import (
	"fmt"
	. "gin.blog.com/admin" //取别名
	api2 "gin.blog.com/api"
	"gin.blog.com/common"
	. "gin.blog.com/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)



//令牌验证
func Authorize() gin.HandlerFunc{
	return func(c *gin.Context){
		username := c.PostForm("username") // 用户名
		ts := c.PostForm("ts") // 时间戳
		token := c.PostForm("token") // 访问令牌

		if strings.ToLower(common.MD5([]byte(username+ts+common.TokenSalt))) == strings.ToLower(token) {
			// 验证通过，会继续访问下一个中间件
			c.Next()
		} else {
			// 验证不通过，不再调用后续的函数处理
			c.Abort()
			c.JSON(http.StatusUnauthorized,gin.H{"message":"访问未授权"})
			// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
			return
		}
	}
}

//跨域处理
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		//放行所有OPTIONS方法
		fmt.Println(method)
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		}
		// 处理请求
		c.Next()
	}
}



func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(Cors())

	//store,_ := redis.NewStore(10,"tcp","localhost:6379","",[]byte("secret"))
	//router.Use(sessions.Sessions("mysession",store))

	router.POST("/admin/login", Login)

	api := router.Group("/api")
	{
		api.GET("/column_list", ColumnList)
		api.GET("/category_list", CategoryList)
		api.GET("/article_list", api2.ArticleList)
	}

	admin := router.Group("/admin/v1")
	//中间件验证
	//admin.Use(Cors())
	admin.Use(Jwt())
	{
		admin.GET("/routes", Routers)
		admin.GET("/info", Info)
		admin.POST("/column_add", ColumnAdd)
		admin.GET("/column_list", ColumnList)
		admin.GET("/category_list", CategoryList)
		admin.POST("/category_add", CategoryAdd)
		admin.GET("/category_info", CategoryInfo)
		admin.POST("/category_edit", CategoryEdit)
		//文章
		admin.GET("/article_list", ArticleList)
		admin.POST("/article_add", ArticleAdd)
		admin.GET("/article_info", ArticleInfo)
		admin.POST("/article_edit", ArticleEdit)
	}
	return router
}



