package data

import "database/sql"

// @title	Borrower.Create
// @description	创建借阅者
// @param	UserName	Borrower.UserName
// @param	Password	Borrower.Password
func (borrower *Borrower) Create() (err error) {
	statement := "INSERT INTO Borrower (username, password) VALUES (?, ?)"
	_, err = Db.Query(statement, borrower.UserName, borrower.Password)
	return
}

// @title	Borrower.RetrieveBorrowerByUserName
// @description	根据用户名搜索用户
func (borrower *Borrower) RetrieveBorrowerByUserName() (err error) {
	statement := `SELECT UID, password, Borrower.cardNO, name, major, BorrowerType.borrowerType
				  FROM Borrower LEFT JOIN Card ON Borrower.cardNO = Card.cardNO 
				      LEFT JOIN BorrowerType ON Card.borrowerType = BorrowerType.borrowerType
				  WHERE username=?`
	// 处理还未领取借阅证的情况
	cardNo := sql.NullString{}
	name := sql.NullString{}
	major := sql.NullString{}
	borrowerType := sql.NullString{}

	err = Db.QueryRow(statement, borrower.UserName).Scan(&borrower.UID, &borrower.Password,
		&cardNo, &name, &major, &borrowerType)
	if cardNo.Valid {
		borrower.Card.CardNO = cardNo.String
		borrower.Card.Name = name.String
		borrower.Card.Major = major.String
		borrower.Card.BorrowerType.BorrowerType = borrowerType.String
	}
	return
}
