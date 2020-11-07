package models

type Article struct {
	Id int `gorm:"id" json:"id"`
	Title string `gorm:"title" json:"title"`	//文章标题
	CreateDate string `gorm:"create_date" json:"create_date"`	//创建时间
	Desc string `gorm:"desc" json:"desc"`	//描述 简述
	CategoryId int `gorm:"category_id" json:"category_id"`	//分类ID
	IsShow int `gorm:"is_show" json:"is_show"`		//是否展示
	Content string `gorm:"content" json:"content"`  //内容
}

type ArticleInfo struct {
	Id int `gorm:"id" json:"id"`
}

var BlogArticleTable = "blog_article"
