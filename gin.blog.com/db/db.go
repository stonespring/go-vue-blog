package db

import (
	"gin.blog.com/config"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	_ "github.com/go-sql-driver/mysql"
)

var engine *gorm.DB

func InitDB() error {
	var err error
	conf := config.GetConfig()
	engine, err = gorm.Open("mysql", conf.DBConfig.DbUser+":"+conf.DBConfig.DbPassword+"@tcp("+conf.DBConfig.DbHost+":"+conf.DBConfig.DbPort+")/"+conf.DBConfig.DbName+"?charset=utf8")
	if err != nil {
		err = errors.Wrap(err, "初始化数据库失败")
		return err
	}
	//defer engine.Close()
	return nil
}



func GetDb() *gorm.DB  {
	return engine
}