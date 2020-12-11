package tools

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"library/config"
)

func ServerConfig() config.Server {
	// 返回服务器相关配置
	var server config.Server
	//读取配置文件
	serverConfig, err := ioutil.ReadFile("../config/config.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(serverConfig, &server)
	if err != nil {
		panic(err)
	}
	return server
}

func SqlConfig() config.Sql {
	// 返回数据库相关配置
	var sql config.Sql
	//读取配置文件
	sqlConfig, err := ioutil.ReadFile("../config/config.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(sqlConfig, &sql)
	if err != nil {
		panic(err)
	}
	return sql
}
