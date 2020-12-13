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

// @title		Administrator.GetPassword
// @description	根据账号获取管理员密码
// @param		Administrator.UserName	string	"用户名"
// @return		err						error	"错误信息"
func (administrator *Administrator) RetrieveByUserName() (err error) {
	statement := "SELECT administratorID, password FROM Administrator WHERE username=?"
	err = Db.QueryRow(statement, administrator.UserName).Scan(&administrator.AdministratorID, &administrator.Password)
	return
}
