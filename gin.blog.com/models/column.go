package models

type Column struct {
	Id int `gorm:"id" json:"id"`
	Name string `gorm:"name" json:"name"`	//名称
	Top string `gorm:"top" json:"top"`		//排序
	IsShow string `gorm:"is_show" json:"is_show"`  //是否展示
}

var BlogColumnTable = "blog_column"
