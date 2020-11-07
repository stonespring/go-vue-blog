package config

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
)

//配置数据库应用json
type DBConfig struct {
	DbHost	string	`json:"db_host"`
	DbUser	string	`json:"db_user"`
	DbPassword string	`json:"db_password"`
	DbPort	string	`json:"db_port"`
	DbName	string	`json:"db_name"`
}


type RedisConfig struct {
	Addr	string `json:"addr"`
	Password string	`json:"password"`
	DB 	int	`json:"db"`
}

//调用json文件
type Config struct {
	DBConfig DBConfig	`json:"db_config"`
	RedisConfig RedisConfig		`json:"redis_config"`
}

var conf Config

func InitConfig(configPath string) error {
	ConfigFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		err = errors.Wrap(err,"获取配置失败")
		return err
	}

	err = json.Unmarshal(ConfigFile, &conf)
	if err != nil {
		err = errors.Wrap(err, "解析json数据失败")
	}
	return nil
}

func GetConfig() Config {
	return conf
}