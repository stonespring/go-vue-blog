package api

import (
	"gin.blog.com/common"
	"gin.blog.com/db"
	"gin.blog.com/models"
	"github.com/gin-gonic/gin"
	"net/http"
)



//文章列表
func ArticleList(c *gin.Context)  {
	var Article []models.Article
	category_id := c.Query("category_id")

	where   := map[string]interface{}{}

	if category_id != "0"  {
		where["category_id"] = category_id
		db.GetDb().Table(models.BlogArticleTable).Where(where).Order("id desc").Find(&Article)
	}else{
		db.GetDb().Debug().Table(models.BlogArticleTable).Find(&Article)
	}


	c.JSON(http.StatusOK, gin.H{
		"code" : common.SUCCESS,
		"msg" : common.GetMsg(common.SUCCESS),
		"data": Article,
	})
}

//编辑文章显示
func ArticleInfo(c *gin.Context){
	where   := map[string]interface{}{}
	formArticle := models.Article{}
	//获取get参数
	id := c.Query("id");
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code" : common.ERROR,
			"msg" : "请求错误",
		})
	}
	where["id"] = id

	res := db.GetDb().Table(models.BlogArticleTable).Where(where).First(&formArticle)
	code := common.ERROR
	msg := common.GetMsg(code)
	if res != nil {
		code = common.SUCCESS
		msg = common.GetMsg(code)
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : msg,
		"data": formArticle,
	})
}

//文章列表
func ArticleNumList(c *gin.Context)  {
	var Article []models.Article
	category_id := c.Query("category_id")
	where   := map[string]interface{}{}
	if category_id != "" {
		where["category_id"] = category_id
		db.GetDb().Table(models.BlogArticleTable).Where(where).Order("id desc").Limit(10).Find(&Article)
	}else{
		db.GetDb().Debug().Table(models.BlogArticleTable).Limit(10).Order("create_date desc").Find(&Article)
	}


	c.JSON(http.StatusOK, gin.H{
		"code" : common.SUCCESS,
		"msg" : common.GetMsg(common.SUCCESS),
		"data": Article,
	})
}