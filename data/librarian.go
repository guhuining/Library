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
	statement := "SELECT librarianID, password FROM Librarian WHERE username=?"
	err = Db.QueryRow(statement, librarian.UserName).Scan(&librarian.LibrarianID, &librarian.Password)
	return
}

// @title		Librarian.Delete
// @description 删除图书管理员
// @param		Librarian.LibrarianID	string	"管理员ID"
// @return		err						error	"错误信息"
func (librarian *Librarian) Delete() (err error) {
	statement := `DELETE FROM Librarian WHERE librarianID=?`
	_, err = Db.Query(statement, librarian.LibrarianID)
	return
}
