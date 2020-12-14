package data

// @title		Librarian.Create
// @description	添加图书管理员
// @param		Librarian.UserName	string	"用户名"
// @param		Librarian.Password	string	"密码"
// @return		err					error	"错误信息"
func (librarian *Librarian) Create() (err error) {
	statement := "INSERT INTO librarian (username, password) VALUES(?, ?)"
	_, err = Db.Query(statement, librarian.UserName, librarian.Password)
	return
}

// @title		Librarian.GetPassword
// @description	根据账号获取图书管理员密码
// @param		Librarian.UserName	string	"用户名"
// @return		err					error	"错误信息"
func (librarian *Librarian) RetrieveByUserName() (err error) {
	statement := "SELECT administratorID, password FROM Administrator WHERE username=?"
	err = Db.QueryRow(statement, librarian.UserName).Scan(&librarian.LibrarianID, &librarian.Password)
	return
}
