package data

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"library/tools"
)

var Db *sql.DB

func init() {
	//初始化数据库连接
	var err error
	sqlConfig := tools.SqlConfig()
	connMsg := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		sqlConfig.UserName, sqlConfig.Password, sqlConfig.Host, sqlConfig.Port, sqlConfig.DBName)
	Db, err = sql.Open("mysql", connMsg)
	if err != nil {
		panic(err)
	}
}
