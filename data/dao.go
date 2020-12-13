package data

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"library/config"
)

var Db *sql.DB

func init() {
	//初始化数据库连接
	var err error
	c := &config.Config{}
	databaseConfig := c.GetDatabase()
	connMsg := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		databaseConfig.UserName, databaseConfig.Password, databaseConfig.Host, databaseConfig.Port, databaseConfig.DBName)
	Db, err = sql.Open("mysql", connMsg)
	if err != nil {
		panic(err)
	}
}
