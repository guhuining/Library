package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Database struct {
	UserName string `yaml:"user_name"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"db_name"`
}

// 返回服务器相关配置
func (config *Config) GetServer() Server {
	serverConfig, err := ioutil.ReadFile("config/config.yml") //读取配置文件
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(serverConfig, &config)
	if err != nil {
		panic(err)
	}
	return config.Server
}

// 返回数据库相关配置
func (config *Config) GetDatabase() Database {
	sqlConfig, err := ioutil.ReadFile("config/config.yml") //读取配置文件
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(sqlConfig, &config)
	if err != nil {
		panic(err)
	}
	return config.Database
}
