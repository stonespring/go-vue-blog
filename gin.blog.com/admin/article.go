package controller

import (
	"gin.blog.com/common"
	"gin.blog.com/db"
	"gin.blog.com/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"time"
)

// 添加栏目
func ColumnAdd(c *gin.Context){
	formColumn := models.Column{}
	c.BindJSON(&formColumn)
	res := db.GetDb().Table(models.BlogColumnTable).Create(&formColumn)
	code := common.ERROR
	msg := common.GetMsg(code)
	if res != nil {
		code = common.SUCCESS
		msg = common.GetMsg(code)
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : msg,
	})
}


// 栏目列表
func ColumnList(c *gin.Context){
	var Columns []models.Column
	db.GetDb().Table(models.BlogColumnTable).Find(&Columns)
	data := make(map[string]interface{})
	data["data"] = Columns
	c.JSON(http.StatusOK, gin.H{
		"code" : common.SUCCESS,
		"msg" : common.GetMsg(common.SUCCESS),
		"data": data,
	})
}
//分类列表
func CategoryList(c *gin.Context)  {
	var Category []models.Category
	db.GetDb().Table(models.BlogCategoryTable).Find(&Category)
	data := make(map[string]interface{})
	data["data"] = Category
	c.JSON(http.StatusOK, gin.H{
		"code" : common.SUCCESS,
		"msg" : common.GetMsg(common.SUCCESS),
		"data": data,
	})
}
//添加分类
func CategoryAdd(c *gin.Context){
	formCategory := models.Category{}
	c.ShouldBind(&formCategory)
	res := db.GetDb().Table(models.BlogCategoryTable).Create(&formCategory)
	code := common.ERROR
	msg := common.GetMsg(code)
	if res != nil {
		code = common.SUCCESS
		msg = common.GetMsg(code)
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : msg,
	})
}

var (
	where   = map[string]interface{}{}
)
//编辑展示分类
func CategoryInfo(c *gin.Context){
	formCategory := models.Category{}
	//获取get参数
	id := c.Query("id");
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code" : common.ERROR,
			"msg" : "请求错误",
		})
	}
	where["id"] = id

	res := db.GetDb().Table(models.BlogCategoryTable).Where(where).First(&formCategory)
	//db.GetDb().Debug().Table(models.BlogCategoryTable).Where(where).First(&formCategory)
	code := common.ERROR
	data := make(map[string]interface{})
	data["data"] = formCategory
	msg := common.GetMsg(code)
	if res != nil {
		code = common.SUCCESS
		msg = common.GetMsg(code)
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : msg,
		"data": data,
	})
}

//编辑提交
func CategoryEdit(c *gin.Context)  {

	editCategoryData := models.Category{}
	c.ShouldBind(&editCategoryData)
	where["id"] = editCategoryData.Id
	data := make(map[string]interface{})
	data["category_name"] = editCategoryData.CategoryName
	data["column_id"] = editCategoryData.ColumnId
	data["is_show"] = editCategoryData.IsShow
	data["top"] = editCategoryData.Top
	res := db.GetDb().Table(models.BlogCategoryTable).Where(where).Update(data)

	code := common.ERROR
	msg := common.GetMsg(code)
	if res != nil {
		code = common.SUCCESS
		msg = "修改成功"
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : msg,
	})
}

func ArticleAdd(c *gin.Context)  {
	formAdd := models.Article{}
	c.ShouldBind(&formAdd)
	layout := "2006-01-02T15:04:05.000Z07:00"
	layoutTwo := "2006-01-02 15:04:05"  //转化所需模板
	t,_ := time.Parse(layout,formAdd.CreateDate)
	fmt.Println(t.Unix())
	formAdd.CreateDate = time.Unix(t.Unix(), 0).Format(layoutTwo)
	fmt.Println(formAdd.CreateDate)
	res := db.GetDb().Table(models.BlogArticleTable).Create(&formAdd)
	code := common.ERROR
	msg := common.GetMsg(code)
	if res != nil {
		code = common.SUCCESS
		msg = "文章添加成功"
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : msg,
	})
}

//文章列表
func ArticleList(c *gin.Context)  {
	var Article []models.Article
	db.GetDb().Table(models.BlogArticleTable).Find(&Article)
	c.JSON(http.StatusOK, gin.H{
		"code" : common.SUCCESS,
		"msg" : common.GetMsg(common.SUCCESS),
		"data": Article,
	})
}

//编辑文章显示
func ArticleInfo(c *gin.Context){
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


func ArticleEdit(c *gin.Context)  {

	editArticleData := models.Article{}
	c.ShouldBind(&editArticleData)
	where["id"] = editArticleData.Id
	data := make(map[string]interface{})
	data["title"] = editArticleData.Title
	data["create_date"] = editArticleData.CreateDate
	data["is_show"] = editArticleData.IsShow
	data["content"] = editArticleData.Content
	data["desc"] = editArticleData.Desc
	data["category_id"] = editArticleData.CategoryId

	res := db.GetDb().Table(models.BlogArticleTable).Where(where).Update(data)

	code := common.ERROR
	msg := common.GetMsg(code)
	if res != nil {
		code = common.SUCCESS
		msg = "修改成功"
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : msg,
	})
}
