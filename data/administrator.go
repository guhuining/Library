package data

// @title		Administrator.Create
// @description	添加系统管理员
// @param		Administrator.UserName	string	"用户名"
// @param		Administrator.Password	string	"密码"
// @return		err						error	"错误信息"
func (administrator *Administrator) Create() (err error) {
	statement := "INSERT INTO administrator (username, password) VALUES(?, ?)"
	_, err = Db.Query(statement, administrator.UserName, administrator.Password)
	return
}
