package data

// @title	Borrower.Create
// @description	创建借阅者
// @param	UserName	Borrower.UserName
// @param	Password	Borrower.Password
func (borrower *Borrower) Create() (err error) {
	statement := "INSERT INTO Borrower (username, password) VALUES (?, ?)"
	_, err = Db.Query(statement, borrower.UserName, borrower.Password)
	return
}
