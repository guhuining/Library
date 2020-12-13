package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Sql struct {
	UserName string `yaml:"user_name"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"db_name"`
}

// 返回服务器相关配置
func (server *Server) ServerConfig() {
	serverConfig, err := ioutil.ReadFile("config.yml") //读取配置文件
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(serverConfig, &server)
	if err != nil {
		panic(err)
	}
}

// 返回数据库相关配置
func (sql *Sql) SqlConfig() {

	sqlConfig, err := ioutil.ReadFile("config.yml") //读取配置文件
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(sqlConfig, &sql)
	if err != nil {
		panic(err)
	}
}
