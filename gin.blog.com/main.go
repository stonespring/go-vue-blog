package main

import (
	"flag"
	"fmt"
	"gin.blog.com/config"
	"gin.blog.com/db"
	"gin.blog.com/redis"
	. "gin.blog.com/router"
	"github.com/gorilla/sessions"
	"os"
)

var session *sessions.Session

func main() {

	var configPath string
	flag.StringVar(&configPath, "config", "./config/config.json", "配置文件路径") // name 等于接收值, value 默认值 usage 提示
	flag.Parse()



	var err error
	err = config.InitConfig(configPath)  //todo 2020-08-20 获取配置文件信息 且解析以
	if err != nil {
		fmt.Printf("初始化配置文件失败,错误信息:%v", err)
		os.Exit(1)
	}

	err = db.InitDB() //todo 初始化数据库
	if err != nil {
		fmt.Printf("数据库连接错误")
	}

	err = redis.InitRedis() //todo 连接redis
	if err != nil {
		fmt.Printf("redis连接错误")
	}

	//路由
	router := InitRouter()
	//当整个程序完成之后关闭数据库连
	defer db.GetDb().Close()
	router.Run(":8080")

}

func Session() *sessions.Session {
	return session
}