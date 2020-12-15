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

// @title Card.Insert
// @description	插入新借阅证
// @param	UID				Borrower.UID				借阅者标识符
// @param	CardNO			Card.CardNO					借阅证号码
// @param	Name			Card.Name					姓名
// @param	Major			Card.Major					专业
// @param	BorrowerType	BorrowerTYpe.BorrowerType	借阅者类型
// @return
func (borrower *Borrower) BindCard() (err error) {
	tx, err := Db.Begin()
	if err != nil {
		return
	}
	statement := `INSERT INTO Card (cardNO, name, major, borrowerType) VALUES (?, ?, ?, ?)`
	_, err = Db.Query(statement, borrower.Card.CardNO, borrower.Card.Name, borrower.Card.Major,
		borrower.Card.BorrowerType.BorrowerType)
	if err != nil {
		tx.Rollback()
		return
	}
	// 在borrower表中更新卡信息
	statement = "UPDATE Borrower SET cardNO = ? WHERE UID = ?"
	_, err = tx.Query(statement, borrower.Card.CardNO, borrower.UID)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}
