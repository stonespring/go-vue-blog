package models

type Category struct {
	Id int `gorm:"id" json:"id"`
	CategoryName string `gorm:"name" json:"category_name"`	//分类名称
	ColumnId int `gorm:"column_id" json:"column_id"`	//栏目ID
	Top string `gorm:"top" json:"top"`		//排序
	IsShow string `gorm:"is_show" json:"is_show"`  //是否展示
}

type CategoryInfo struct {
	Id int `gorm:"id" json:"id"`
}

var BlogCategoryTable = "blog_category"
