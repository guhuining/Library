package data

import (
	"database/sql"
	"library/my_error"
	"time"
)

// @title BorrowItem.Borrow
// @description	借书
// @param	CardNO			BorrowItem.CardNO			借阅证号码
// @param	PublicationID	Publication.PublicationID	出版物ID
func (borrowItem *BorrowItem) Borrow() (err error) {
	tx, err := Db.Begin()
	if err != nil {
		return
	}
	// 检查是否有未归还图书
	statement := `SELECT COUNT(*) FROM BorrowItem JOIN Card ON BorrowItem.cardNO = Card.cardNO
				  								  JOIN BorrowerType ON Card.borrowerType = BorrowerType.borrowerType
                  WHERE BorrowItem.cardNO = ? AND DATE_ADD(borrowDate, INTERVAL BorrowerType.period DAY) < now()
                  AND status = 0`
	rows, err := tx.Query(statement, borrowItem.Card.CardNO)
	if err != nil {
		tx.Rollback()
		return
	}
	if rows.Next() {
		var count = 0
		err = rows.Scan(&count)
		if err != nil {
			tx.Rollback()
			return
		} else if count != 0 {
			err = my_error.BorrowOutOfTimeError
			tx.Rollback()
			return
		}
	}
	rows.Close()
	// 库存-1
	statement = `UPDATE Publication SET inventory = inventory - 1 WHERE PublicationID = ?`
	_, err = tx.Query(statement, borrowItem.Publication.PublicationID)
	if err != nil {
		tx.Rollback()
		err = my_error.InventoryNotEnoughError
		return
	}
	// 添加借阅信息
	statement = `INSERT INTO BorrowItem (cardNO, publicationID, borrowDate) VALUES(?, ?, NOW())`
	_, err = tx.Query(statement, borrowItem.Card.CardNO, borrowItem.Publication.PublicationID)
	if err != nil {
		tx.Rollback()
		return
	}
	// 查询借阅量是否达到上限
	statement = `SELECT currentBorrowNumber FROM Card INNER JOIN BorrowerType ON Card.borrowerType = BorrowerType.borrowerType 
				 WHERE cardNO = ? AND currentBorrowNumber < maxBorrowNumber`
	err = tx.QueryRow(statement, borrowItem.Card.CardNO).Scan(&borrowItem.Card.CurrentBorrowNumber)
	if err != nil && err.Error() == sql.ErrNoRows.Error() {
		err = my_error.MaxBorrowNumberError
		tx.Rollback()
		return
	} else if err != nil {
		return
	}
	statement = `UPDATE Card INNER JOIN BorrowerType ON Card.borrowerType = BorrowerType.borrowerType 
				 SET currentBorrowNumber = currentBorrowNumber + 1 WHERE cardNO = ? 
				 AND currentBorrowNumber < maxBorrowNumber`
	_, err = tx.Query(statement, borrowItem.Card.CardNO)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// @title BorrowItem.IsOutOfTime
// @description	查询是否逾期
// @param	borrowItemID	BorrowItem.borrowItemID	借阅订单ID
func (borrowItem *BorrowItem) IsOutOfTime() (err error) {
	statement := `SELECT COUNT(*) FROM BorrowItem JOIN Card ON BorrowItem.cardNO = Card.cardNO
				  								  JOIN BorrowerType ON Card.borrowerType = BorrowerType.borrowerType
                  WHERE BorrowItem.borrowItemID = ? AND DATE_ADD(borrowDate, INTERVAL BorrowerType.period DAY) < now()`
	rows, err := Db.Query(statement, borrowItem.BorrowItemID)
	if err != nil {
		return
	}
	defer rows.Close()
	if rows.Next() {
		var count = 0
		err = rows.Scan(&count)
		if err != nil {
			return
		} else if count != 0 {
			err = my_error.BorrowOutOfTimeError
			return
		}
	}
	return
}

// @title BorrowItem.GetFine
// @description	获取超时罚金
// @param	publicationID	Publication.PublicationID	出版物ID
func (borrowItem *BorrowItem) GetFine() (err error) {
	statement := `SELECT fine FROM BorrowItem
    			  JOIN Publication ON BorrowItem.publicationID = Publication.publicationID
    			  JOIN PublicationType ON Publication.publicationType = PublicationType.publicationType
    			  WHERE BorrowItem.borrowItemID = ?`
	err = Db.QueryRow(statement, borrowItem.BorrowItemID).Scan(&borrowItem.Publication.PublicationType.Fine)
	return
}

// @title BorrowItem.Return
// @description	还书
// @param	BorrowItemID	BorrowItem.BorrowItemID	借阅订单ID
func (borrowItem *BorrowItem) Return() (err error) {
	statement := `UPDATE BorrowItem JOIN Card ON BorrowItem.cardNO = Card.cardNO 
				  JOIN Publication ON BorrowItem.publicationID = Publication.publicationID
				  SET status = 1, Card.currentBorrowNumber = Card.currentBorrowNumber - 1, dueDate = NOW(), 
				      Publication.inventory=Publication.inventory+1
				  WHERE BorrowItem.BorrowItemID = ? AND status = 0`
	_, err = Db.Query(statement, borrowItem.BorrowItemID)
	return
}

// @title BorrowItem.GetBorrowItem
func (borrowItem *BorrowItem) GetBorrowItem() (results []BorrowItem, err error) {
	statement := `SELECT b.borrowItemID, p.name, p.author, b.borrowDate
				  FROM BorrowItem b JOIN publication p on b.publicationID = p.publicationID
				  WHERE b.cardNO = ? AND b.status = 0`
	rows, err := Db.Query(statement, borrowItem.Card.CardNO)
	if err != nil {
		return
	}
	for rows.Next() {
		var temp BorrowItem
		var t string
		err = rows.Scan(&temp.BorrowItemID, &temp.Publication.Name, &temp.Publication.Author, &t)
		temp.BorrowDate, _ = time.Parse("2006-01-02 15:04:05", t)
		if err != nil {
			return
		}
		results = append(results, temp)
	}
	return
}

func (borrowItem *BorrowItem) GetPrice() (err error) {
	statement := `SELECT price FROM BorrowItem b 
    			  JOIN Publication p ON b.publicationID=p.publicationID
    			  WHERE b.borrowItemID=?`
	err = Db.QueryRow(statement, borrowItem.BorrowItemID).Scan(&borrowItem.Publication.Price)
	return
}
